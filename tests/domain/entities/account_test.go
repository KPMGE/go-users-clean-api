package entities_test

import (
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	fakeAccountUserName string = "any_name"
	fakeAccountEmail    string = "any_valid_email@gmail.com"
	fakeAccountPassword string = "any_password"
)

func TestNewAccount_WithRightData(t *testing.T) {
	account, err := entities.NewAccount(fakeAccountUserName, fakeAccountEmail, fakeAccountPassword)

	require.Nil(t, err)
	require.Equal(t, account.Email, fakeAccountEmail)
	require.Equal(t, account.UserName, fakeAccountUserName)
	require.Equal(t, account.Password, fakeAccountPassword)
}

func TestNewAccount_WithInvalidEmail(t *testing.T) {
	account, err := entities.NewAccount(fakeAccountUserName, "any_invalid_email", fakeAccountPassword)

	require.Nil(t, account)
	require.Error(t, err)
	require.Equal(t, err.Error(), "Invalid email!")
}

func TestNewAccount_WithBlankFields(t *testing.T) {
	account, err := entities.NewAccount("", fakeAccountEmail, fakeAccountPassword)
	require.Nil(t, account)
	require.Error(t, err)

	account, err = entities.NewAccount(fakeAccountUserName, "", fakeAccountPassword)
	require.Nil(t, account)
	require.Error(t, err)

	account, err = entities.NewAccount(fakeAccountUserName, fakeAccountEmail, "")
	require.Nil(t, account)
	require.Error(t, err)
}
