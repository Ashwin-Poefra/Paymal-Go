package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	// 1. Hashing a password should return no errors
	password := RandomString(6)
	hashedPassword1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	// 2. A password match should have no errors
	err = CheckPassword(password, hashedPassword1)
	require.NoError(t, err)

	// 3. A password mismatch should have an error
	wrongPassword := RandomString(7)
	err = CheckPassword(wrongPassword, hashedPassword1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	// 4. A password hashed twice should return 2 different hashed passwords
	hashedPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}
