package db

type TargetAccountJSON struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	AtmBankCode     string `json:"atmBankCode"`
	BankDetail      int64  `json:"bankDetail"`
	AccountType     string `json:"accountType"`
	AccountNumber   string `json:"accountNumber"`
	Bank            string `json:"bank"`
	BankBranch      string `json:"bankBranch"`
	Description     string `json:"description"`
	TargetType      int64  `json:"targetType"`
	Status          int64  `json:"status"`
	FourthDigit     string `json:"fourthDigit"`
	CustomerID      int64  `json:"customerId"`
	Currency        string `json:"currency"`
	AccountTypeCode string `json:"accountTypeCode"`
	Amount          string `json:"amount"`
	IsFavorite      string `json:"isFavorite"`
}

func ConverTargetAccountModelToJSON(ta TargetAccount) TargetAccountJSON {
	return TargetAccountJSON{
		ID:              ta.ID,
		Name:            ta.Name.String,
		AtmBankCode:     ta.AtmBankCode.String,
		BankDetail:      ta.BankDetail.Int64,
		AccountType:     ta.AccountType,
		AccountNumber:   ta.AccountNumber.String,
		Bank:            ta.Bank,
		BankBranch:      ta.BankBranch.String,
		Description:     ta.Description.String,
		TargetType:      ta.TargetType,
		Status:          ta.Status,
		FourthDigit:     ta.FourthDigit.String,
		CustomerID:      ta.CustomerID,
		Currency:        ta.Currency.String,
		AccountTypeCode: ta.AccountTypeCode.String,
		Amount:          ta.Amount.String,
		IsFavorite:      ta.IsFavorite.String,
	}
}
