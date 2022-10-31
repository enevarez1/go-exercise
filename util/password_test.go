package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := "random"

	hashed, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashed)

	err = CheckPassword(password, hashed)
	require.NoError(t, err)

	wrongPassword := "wrongp"
	err = CheckPassword(wrongPassword, hashed)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashed2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashed2)
	require.NotEqual(t, hashed, hashed2)
}