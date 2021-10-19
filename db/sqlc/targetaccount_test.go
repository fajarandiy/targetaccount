package db

import (
	"context"
	"database/sql"
	"testing"

	"com.example/targetaccount/util"
	"github.com/stretchr/testify/require"
)

func TestCreateTargetAccount(t *testing.T) {
	createRandomTargetAccount(t)
}

func createRandomTargetAccount(t *testing.T) TargetAccount {
	arg := CreateTargetAccountParams{
		Name:            sql.NullString{String: util.RandomName(), Valid: true},
		AtmBankCode:     sql.NullString{String: util.RandomAtmBankCode(), Valid: true},
		BankDetail:      sql.NullInt64{Int64: util.RandomBankDetail(), Valid: true},
		AccountType:     util.RandomAccountType(),
		AccountNumber:   sql.NullString{String: util.RandomAccountNumber(), Valid: true},
		Bank:            util.RandomBank(),
		Description:     sql.NullString{String: util.RandomDescription(), Valid: true},
		TargetType:      util.RandomTargetType(),
		Status:          27, // Active
		FourthDigit:     sql.NullString{String: util.RandomFourthDigit(), Valid: true},
		CustomerID:      util.RandomCustomerID(),
		Currency:        sql.NullString{String: util.RandomCurrency(), Valid: true},
		AccountTypeCode: sql.NullString{String: util.RandomAccountTypeCode(), Valid: true},
		Amount:          sql.NullString{String: util.RandomAmount(), Valid: true},
		IsFavorite:      sql.NullString{String: util.RandomIsFavorite(), Valid: true},
	}

	var err error
	var result sql.Result

	result, err = testQueries.CreateTargetAccount(context.Background(), arg)
	require.NoError(t, err)

	var lastInsertedId int64
	lastInsertedId, err = result.LastInsertId()
	require.NoError(t, err)
	require.NotEmpty(t, lastInsertedId)

	var targetAccount TargetAccount
	targetAccount, err = testQueries.GetTargetAccount(context.Background(), lastInsertedId)
	require.NoError(t, err)
	require.NotEmpty(t, targetAccount)
	require.NotZero(t, targetAccount.ID)
	require.Equal(t, lastInsertedId, targetAccount.ID)
	require.Equal(t, arg.Name, targetAccount.Name)
	require.Equal(t, arg.AtmBankCode, targetAccount.AtmBankCode)
	require.Equal(t, arg.BankDetail, targetAccount.BankDetail)
	require.Equal(t, arg.AccountType, targetAccount.AccountType)
	require.Equal(t, arg.AccountNumber, targetAccount.AccountNumber)
	require.Equal(t, arg.Bank, targetAccount.Bank)
	require.Equal(t, arg.Description, targetAccount.Description)
	require.Equal(t, arg.TargetType, targetAccount.TargetType)
	require.Equal(t, arg.Status, targetAccount.Status)
	require.Equal(t, arg.FourthDigit, targetAccount.FourthDigit)
	require.Equal(t, arg.CustomerID, targetAccount.CustomerID)
	require.Equal(t, arg.Currency, targetAccount.Currency)
	require.Equal(t, arg.AccountTypeCode, targetAccount.AccountTypeCode)
	// require.Equal(t, arg.Amount, targetAccount.Amount)
	require.Equal(t, arg.IsFavorite, targetAccount.IsFavorite)

	return targetAccount
}

func TestDeleteTargetAccount(t *testing.T) {
	createdTargetAccount := createRandomTargetAccount(t)

	arg := DeleteTargetAccountParams{
		ID:     createdTargetAccount.ID,
		Status: 29, // Inactive
	}

	var err error
	var result sql.Result

	result, err = testQueries.DeleteTargetAccount(context.Background(), arg)
	require.NoError(t, err)

	var rowsAffected int64
	rowsAffected, err = result.RowsAffected()
	require.NoError(t, err)
	require.NotZero(t, rowsAffected)

	var inactivatedTargetAccount TargetAccount
	inactivatedTargetAccount, err = testQueries.GetTargetAccount(context.Background(), arg.ID)
	require.NoError(t, err)
	require.Equal(t, arg.Status, inactivatedTargetAccount.Status)
}

func TestGetTargetAccount(t *testing.T) {
	createdTargetAccount := createRandomTargetAccount(t)
	tragetAccount, err := testQueries.GetTargetAccount(context.Background(), createdTargetAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tragetAccount)

	require.Equal(t, createdTargetAccount.ID, tragetAccount.ID)
	require.Equal(t, createdTargetAccount.Name, tragetAccount.Name)
	require.Equal(t, createdTargetAccount.AtmBankCode, tragetAccount.AtmBankCode)
	require.Equal(t, createdTargetAccount.BankDetail, tragetAccount.BankDetail)
	require.Equal(t, createdTargetAccount.AccountType, tragetAccount.AccountType)
	require.Equal(t, createdTargetAccount.AccountNumber, tragetAccount.AccountNumber)
	require.Equal(t, createdTargetAccount.Bank, tragetAccount.Bank)
	require.Equal(t, createdTargetAccount.Description, tragetAccount.Description)
	require.Equal(t, createdTargetAccount.TargetType, tragetAccount.TargetType)
	require.Equal(t, createdTargetAccount.Status, tragetAccount.Status)
	require.Equal(t, createdTargetAccount.FourthDigit, tragetAccount.FourthDigit)
	require.Equal(t, createdTargetAccount.CustomerID, tragetAccount.CustomerID)
	require.Equal(t, createdTargetAccount.Currency, tragetAccount.Currency)
	require.Equal(t, createdTargetAccount.AccountTypeCode, tragetAccount.AccountTypeCode)
	require.Equal(t, createdTargetAccount.Amount, tragetAccount.Amount)
	require.Equal(t, createdTargetAccount.IsFavorite, tragetAccount.IsFavorite)
}

func TestGetTargetAccountForUpdate(t *testing.T) {
	createdTargetAccount := createRandomTargetAccount(t)
	tragetAccount, err := testQueries.GetTargetAccountForUpdate(context.Background(), createdTargetAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tragetAccount)

	require.Equal(t, createdTargetAccount.ID, tragetAccount.ID)
	require.Equal(t, createdTargetAccount.Name, tragetAccount.Name)
	require.Equal(t, createdTargetAccount.AtmBankCode, tragetAccount.AtmBankCode)
	require.Equal(t, createdTargetAccount.BankDetail, tragetAccount.BankDetail)
	require.Equal(t, createdTargetAccount.AccountType, tragetAccount.AccountType)
	require.Equal(t, createdTargetAccount.AccountNumber, tragetAccount.AccountNumber)
	require.Equal(t, createdTargetAccount.Bank, tragetAccount.Bank)
	require.Equal(t, createdTargetAccount.Description, tragetAccount.Description)
	require.Equal(t, createdTargetAccount.TargetType, tragetAccount.TargetType)
	require.Equal(t, createdTargetAccount.Status, tragetAccount.Status)
	require.Equal(t, createdTargetAccount.FourthDigit, tragetAccount.FourthDigit)
	require.Equal(t, createdTargetAccount.CustomerID, tragetAccount.CustomerID)
	require.Equal(t, createdTargetAccount.Currency, tragetAccount.Currency)
	require.Equal(t, createdTargetAccount.AccountTypeCode, tragetAccount.AccountTypeCode)
	require.Equal(t, createdTargetAccount.Amount, tragetAccount.Amount)
	require.Equal(t, createdTargetAccount.IsFavorite, tragetAccount.IsFavorite)
}

func TestListTargetAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTargetAccount(t)
	}

	arg := ListTargetAccountParams{
		Limit:  5,
		Offset: 5,
	}

	listTargetAccount, err := testQueries.ListTargetAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, listTargetAccount)
	require.Len(t, listTargetAccount, 5)

	for _, targetAccount := range listTargetAccount {
		require.NotEmpty(t, targetAccount)
	}
}

func TestUpdateTargetAccount(t *testing.T) {
	createdTargetAccount := createRandomTargetAccount(t)

	arg := UpdateTargetAccountParams{
		ID:          createdTargetAccount.ID,
		Description: sql.NullString{String: "Test update target account", Valid: true},
	}

	var err error
	var result sql.Result

	result, err = testQueries.UpdateTargetAccount(context.Background(), arg)
	require.NoError(t, err)

	var rowsAffected int64
	rowsAffected, err = result.RowsAffected()
	require.NoError(t, err)
	require.NotZero(t, rowsAffected)

	var updatedTargetAccount TargetAccount
	updatedTargetAccount, err = testQueries.GetTargetAccount(context.Background(), arg.ID)
	require.NoError(t, err)
	require.Equal(t, arg.Description, updatedTargetAccount.Description)
}
