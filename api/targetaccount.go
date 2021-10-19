package api

import (
	"database/sql"
	"net/http"

	db "com.example/targetaccount/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createTargetAccountRequest struct {
	Name            string `json:"name" binding:"required"`
	AtmBankCode     string `json:"atmBankCode"`
	BankDetail      int64  `json:"bankDetail" binding:"required"`
	AccountType     string `json:"accountType" binding:"required"`
	AccountNumber   string `json:"accountNumber"`
	Bank            string `json:"bank"`
	BankBranch      string `json:"bankBranch"`
	Description     string `json:"description"`
	TargetType      int64  `json:"targetType" binding:"required"`
	FourthDigit     string `json:"fourthDigit"`
	CustomerID      int64  `json:"customerId" binding:"required"`
	Currency        string `json:"currency"`
	AccountTypeCode string `json:"accountTypeCode"`
	Amount          string `json:"amount" binding:"numeric"`
	IsFavorite      string `json:"isFavorite"`
}

func (server *Server) createTargetAccount(ctx *gin.Context) {
	var req createTargetAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateTargetAccountParams{
		Name:            sql.NullString{String: req.Name, Valid: true},
		AtmBankCode:     sql.NullString{String: req.AtmBankCode, Valid: true},
		BankDetail:      sql.NullInt64{Int64: req.BankDetail, Valid: true},
		AccountType:     req.AccountType,
		AccountNumber:   sql.NullString{String: req.AccountNumber, Valid: true},
		Bank:            req.Bank,
		BankBranch:      sql.NullString{String: req.BankBranch, Valid: true},
		Description:     sql.NullString{String: req.Description, Valid: true},
		TargetType:      req.TargetType,
		Status:          27, // Active
		FourthDigit:     sql.NullString{String: req.FourthDigit, Valid: true},
		CustomerID:      req.CustomerID,
		Currency:        sql.NullString{String: req.Currency, Valid: true},
		AccountTypeCode: sql.NullString{String: req.AccountTypeCode, Valid: true},
		Amount:          sql.NullString{String: req.Amount, Valid: true},
		IsFavorite:      sql.NullString{String: req.IsFavorite, Valid: true},
	}

	result, err := server.store.CreateTargetAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	targetAccount, err := server.store.GetTargetAccount(ctx, lastInsertedId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, db.ConverTargetAccountModelToJSON(targetAccount))
}

type getTargetAccountParams struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getTargetAccount(ctx *gin.Context) {
	var req getTargetAccountParams
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	targetAccount, err := server.store.GetTargetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, db.ConverTargetAccountModelToJSON(targetAccount))
}

type listTargetAccountParams struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) getListTargetAccount(ctx *gin.Context) {
	var req listTargetAccountParams
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListTargetAccountParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	listTargetAccount, err := server.store.ListTargetAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var convertedListTargetAccount []db.TargetAccountJSON
	for _, targetAccount := range listTargetAccount {
		convertedListTargetAccount = append(convertedListTargetAccount, db.ConverTargetAccountModelToJSON(targetAccount))
	}
	ctx.JSON(http.StatusOK, convertedListTargetAccount)
}

type updateTargetAccountRequest struct {
	ID          int64  `json:"id" binding:"required,min=1"`
	Description string `json:"description"`
}

func (server *Server) updateTargetAccount(ctx *gin.Context) {
	var req updateTargetAccountRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err := server.store.GetTargetAccountForUpdate(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	_, err = server.store.UpdateTargetAccount(ctx, db.UpdateTargetAccountParams{
		ID:          req.ID,
		Description: sql.NullString{String: req.Description, Valid: true},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	updatedTarget, err := server.store.GetTargetAccount(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, db.ConverTargetAccountModelToJSON(updatedTarget))
}

type deleteTargetAccountRequest struct {
	ID int64 `json:"id" binding:"required,min=1"`
}

func (server *Server) deleteTargetAccount(ctx *gin.Context) {
	var req deleteTargetAccountRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err := server.store.GetTargetAccountForUpdate(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	_, err = server.store.DeleteTargetAccount(ctx, db.DeleteTargetAccountParams{
		ID:     req.ID,
		Status: 29, // Inactive
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	deletedTargetAccount, err := server.store.GetTargetAccount(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, db.ConverTargetAccountModelToJSON(deletedTargetAccount))
}
