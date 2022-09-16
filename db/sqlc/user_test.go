package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/transparentideas/gobank/util"
)

func createRandomUser(t *testing.T) User {
	password := util.RandomString(10)
	hashedpass, err := util.HashPassword(password)
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:      util.RandomOwner(),
		HashedPasswor: hashedpass,
		FullName:      util.RandomOwner(),
		Email:         util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPasswor, user.HashedPasswor)
	require.Equal(t, arg.FullName, user.FullName)
	require.True(t, user.PasswordChangedAt.IsZero())

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomAccount(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user.Username, user2.Username)
	require.Equal(t, user.FullName, user2.FullName)
	require.Equal(t, user.Email, user2.Email)
	require.Equal(t, user.HashedPasswor, user2.HashedPasswor)
	require.WithinDuration(t, user.CreatedAt, user2.CreatedAt, 1*time.Second)
}
