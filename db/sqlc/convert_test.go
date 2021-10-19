package db

import (
	"database/sql"
	"testing"

	"com.example/targetaccount/util"
	"github.com/stretchr/testify/require"
)

func TestConverTargetAccountModelToResponseAPI(t *testing.T) {
	var targetAccount = TargetAccount{
		ID:              1,
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

	convertedTargetAccount := ConverTargetAccountModelToJSON(targetAccount)

	require.Equal(t, targetAccount.ID, convertedTargetAccount.ID)
	require.Equal(t, targetAccount.Name.String, convertedTargetAccount.Name)
	require.Equal(t, targetAccount.AtmBankCode.String, convertedTargetAccount.AtmBankCode)
	require.Equal(t, targetAccount.BankDetail.Int64, convertedTargetAccount.BankDetail)
	require.Equal(t, targetAccount.AccountType, convertedTargetAccount.AccountType)
	require.Equal(t, targetAccount.AccountNumber.String, convertedTargetAccount.AccountNumber)
	require.Equal(t, targetAccount.Bank, convertedTargetAccount.Bank)
	require.Equal(t, targetAccount.Description.String, convertedTargetAccount.Description)
	require.Equal(t, targetAccount.TargetType, convertedTargetAccount.TargetType)
	require.Equal(t, targetAccount.Status, convertedTargetAccount.Status)
	require.Equal(t, targetAccount.FourthDigit.String, convertedTargetAccount.FourthDigit)
	require.Equal(t, targetAccount.CustomerID, convertedTargetAccount.CustomerID)
	require.Equal(t, targetAccount.Currency.String, convertedTargetAccount.Currency)
	require.Equal(t, targetAccount.AccountTypeCode.String, convertedTargetAccount.AccountTypeCode)
	require.Equal(t, targetAccount.Amount.String, convertedTargetAccount.Amount)
	require.Equal(t, targetAccount.IsFavorite.String, convertedTargetAccount.IsFavorite)
}
