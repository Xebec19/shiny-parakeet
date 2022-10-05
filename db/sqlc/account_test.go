package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Xebec19/shiny-parakeet/util"
	"github.com/stretchr/testify/require"
)

func TestAccount(t *testing.T) {
	user := randomUser()
	args1 := CreateAccountParams{
		AccountName: util.RandomString(10),
		Dob:         time.Now(),
		Address:     sql.NullString{String: util.RandomString(20), Valid: true},
		Description: sql.NullString{String: util.RandomString(20), Valid: true},
		CreatedBy:   user.UserID,
	}
	// test creating account
	account1, err := TestQueries.CreateAccount(context.Background(), args1)
	require.NoError(t, err)
	require.NotEmpty(t, account1)

	args2 := ReadAccountParams{
		AccountID: account1,
		CreatedBy: user.UserID,
	}
	// test read account
	account2, err := TestQueries.ReadAccount(context.Background(), args2)
	require.NoError(t, err)
	require.Equal(t, args1.AccountName, account2.AccountName)
	require.Equal(t, args1.Description, account2.Description)
	require.Equal(t, args1.Address, account2.Address)

	args3 := CreateAccountParams{
		AccountName: util.RandomString(10),
		Dob:         time.Now(),
		Address:     sql.NullString{String: util.RandomString(20), Valid: true},
		Description: sql.NullString{String: util.RandomString(20), Valid: true},
		CreatedBy:   user.UserID,
	}
	// test creating account
	account3, err := TestQueries.CreateAccount(context.Background(), args3)
	require.NoError(t, err)
	require.NotEmpty(t, account3)

	args4 := ReadAccountParams{
		AccountID: account3,
		CreatedBy: user.UserID,
	}
	account4, err := TestQueries.ReadAccount(context.Background(), args4)
	require.NoError(t, err)
	require.NotEqual(t, account2.AccountID, account4.AccountID)
	require.NotEqual(t, account2.AccountName, account4.AccountName)
	require.NotEqual(t, account2.Description, account4.Description)
	require.NotEqual(t, account2.Address, account4.Address)

	args5 := ReadAllAccountParams{
		CreatedBy: user.UserID,
		Offset:    0,
		Limit:     10,
	}
	// test read all accounts
	accounts, err := TestQueries.ReadAllAccount(context.Background(), args5)
	require.NoError(t, err)
	require.Len(t, accounts, 2)
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}

	args6 := DeleteAccountParams{
		CreatedBy: user.UserID,
		AccountID: account4.AccountID,
	}
	// test delete account
	TestQueries.DeleteAccount(context.Background(), args6)
	args7 := ReadAccountParams{
		AccountID: args6.AccountID,
		CreatedBy: user.UserID,
	}
	account5, err := TestQueries.ReadAccount(context.Background(), args7)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account5)
}
