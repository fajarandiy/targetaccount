package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "com.example/targetaccount/db/mock"
	db "com.example/targetaccount/db/sqlc"
	"com.example/targetaccount/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetTargetAccountAPI(t *testing.T) {
	targetAccount := randomTargetAccount()

	testCases := []struct {
		name            string
		targetAccountID int64
		buildStubs      func(store *mockdb.MockStore)
		checkResponse   func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:            "OK",
			targetAccountID: targetAccount.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTargetAccount(gomock.Any(), gomock.Eq(targetAccount.ID)).
					Times(1).
					Return(targetAccount, nil)
			},
			checkResponse: func(testing *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchTargetAccount(t, recorder.Body, targetAccount)
			},
		},
		{
			name:            "NotFound",
			targetAccountID: targetAccount.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTargetAccount(gomock.Any(), gomock.Eq(targetAccount.ID)).
					Times(1).
					Return(db.TargetAccount{}, sql.ErrNoRows)
			},
			checkResponse: func(testing *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:            "InvalidParameter",
			targetAccountID: 0,
			buildStubs: func(store *mockdb.MockStore) {
			},
			checkResponse: func(testing *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:            "InternalServerError",
			targetAccountID: targetAccount.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTargetAccount(gomock.Any(), gomock.Eq(targetAccount.ID)).
					Times(1).
					Return(db.TargetAccount{}, sql.ErrConnDone)
			},
			checkResponse: func(testing *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// Start test server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			uri := fmt.Sprintf("/targetaccount/%d", tc.targetAccountID)
			request, err := http.NewRequest(http.MethodGet, uri, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})

	}

}

func randomTargetAccount() db.TargetAccount {
	return db.TargetAccount{
		ID:              util.RandomInt(1, 1000),
		Name:            sql.NullString{String: util.RandomName(), Valid: true},
		AtmBankCode:     sql.NullString{String: util.RandomAtmBankCode(), Valid: true},
		BankDetail:      sql.NullInt64{Int64: util.RandomBankDetail(), Valid: true},
		AccountType:     util.RandomAccountType(),
		AccountNumber:   sql.NullString{String: util.RandomAccountNumber(), Valid: true},
		Bank:            util.RandomBank(),
		Description:     sql.NullString{String: util.RandomDescription(), Valid: true},
		TargetType:      util.RandomTargetType(),
		Status:          util.RandomStatus(),
		FourthDigit:     sql.NullString{String: util.RandomFourthDigit(), Valid: true},
		CustomerID:      util.RandomCustomerID(),
		Currency:        sql.NullString{String: util.RandomCurrency(), Valid: true},
		AccountTypeCode: sql.NullString{String: util.RandomAccountTypeCode(), Valid: true},
		Amount:          sql.NullString{String: util.RandomAmount(), Valid: true},
		IsFavorite:      sql.NullString{String: util.RandomIsFavorite(), Valid: true},
	}
}

func requireBodyMatchTargetAccount(t *testing.T, body *bytes.Buffer, targetAccount db.TargetAccount) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var targetAccountJSON db.TargetAccountJSON
	err = json.Unmarshal(data, &targetAccountJSON)
	require.NoError(t, err)
	require.Equal(t, db.ConverTargetAccountModelToJSON(targetAccount), targetAccountJSON)
}
