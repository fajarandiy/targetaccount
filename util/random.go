package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const number = "1234567890"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomStringNumber generates a random string of number
func RandomStringNumber(n int) string {
	var sb strings.Builder
	k := len(number)

	for i := 0; i < n; i++ {
		c := number[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomName generates a random name
func RandomName() string {
	return RandomString(6)
}

// RandomAtmBankCode generates a random ATM bank code
func RandomAtmBankCode() string {
	return RandomStringNumber(3)
}

// RandomBankDetail generates a random id of bank
func RandomBankDetail() int64 {
	return RandomInt(1, 10)
}

// RandomAccountType generates a random account type
func RandomAccountType() string {
	accountType := []string{"Saving Account", "Emoney Account"}
	n := len(accountType)
	return accountType[rand.Intn(n)]
}

// RandomAccountNumber generates a random account number
func RandomAccountNumber() string {
	return RandomStringNumber(10)
}

// RandomBank generates a random bank
func RandomBank() string {
	return "PT. Bank " + RandomString(5)
}

// RandomDescription generates a random description
func RandomDescription() string {
	return RandomString(3)
}

// RandomTargetType generates a random id of target account type
// 9007	: QR Cash Out Withdrawal
// 2922	: Cardless Withdrawal Transfer
// 24 	: INBANK
// 888	: ATMNETWORK
// 23	: Payment
// 8086	: REMMITANCE
// 53	: RTGS
// 205	: SKN
// 2059	: Uangku Transfer
// 231	: Transfer to Virtual Account
func RandomTargetType() int64 {
	targetType := []int64{9007, 2922, 24, 888, 23, 8086, 53, 205, 2059, 231}
	n := len(targetType)
	return targetType[rand.Intn(n)]
}

// RandomStatus generates a random id of target account status
// 27	: Active
// 29	: Inactive
func RandomStatus() int64 {
	status := []int64{27, 29}
	n := len(status)
	return status[rand.Intn(n)]
}

// RandomFourthDigit generates a random fouth digit
func RandomFourthDigit() string {
	return RandomStringNumber(1)
}

// RandomCustomerID generates a random id of customer
func RandomCustomerID() int64 {
	return RandomInt(1, 10)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"IDR", "USD", "EUR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomAccountTypeCode generates a random account type code
func RandomAccountTypeCode() string {
	return RandomStringNumber(4)
}

// RandomAmount generates a random amount
func RandomAmount() string {
	return RandomStringNumber(5)
}

// RandomIsFavorite generates a random is favorite flag
func RandomIsFavorite() string {
	isFavorite := []string{"yes", "no"}
	n := len(isFavorite)
	return isFavorite[rand.Intn(n)]
}
