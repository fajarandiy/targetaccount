package db

import (
	"context"
	"database/sql"
	"testing"

	"com.example/targetaccount/util"
	"github.com/stretchr/testify/require"
)

func TestExecTx(t *testing.T) {
	ctx := context.Background()
	store := NewStore(testDB)
	err := store.ExecTx(ctx, func(q *Queries) error {
		var err error
		var result sql.Result
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

		result, err = q.CreateTargetAccount(ctx, arg)
		if err != nil {
			return err
		}

		var lastInsertedId int64
		lastInsertedId, err = result.LastInsertId()
		if err != nil {
			return err
		}
		require.NotEmpty(t, lastInsertedId)

		var targetAccount TargetAccount
		targetAccount, err = q.GetTargetAccount(context.Background(), lastInsertedId)
		if err != nil {
			return err
		}
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
		return nil
	})
	require.NoError(t, err)
}
