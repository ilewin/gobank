package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(10)
	hash, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hash)
	require.Equal(t, len(hash), 60)
	err = CheckPassword(password, hash)
	require.NoError(t, err)
	wrongPassword := RandomString(10)
	err = CheckPassword(wrongPassword, hash)
	require.Error(t, err)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
	hash2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hash)
	require.NotEqual(t, hash, hash2)
}
