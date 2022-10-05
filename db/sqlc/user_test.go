package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Xebec19/shiny-parakeet/util"
	"github.com/stretchr/testify/require"
)

func randomUser() User {
	args := CreateUserParams{
		FirstName: util.RandomString(10),
		LastName:  sql.NullString{util.RandomString(10), true},
		Email:     util.RandomEmail(),
		Password:  util.RandomString(20),
	}
	user, _ := TestQueries.CreateUser(context.Background(), args)
	return user
}

func TestCreateUser(t *testing.T) {
	user := randomUser()
	require.NotEmpty(t, user)
}
