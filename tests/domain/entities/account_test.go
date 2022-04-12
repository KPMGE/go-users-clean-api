package entities_test

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type Account struct {
	UserName string
	Email    string
	Password string
}

const (
	fakeAccountUserName string = "any_name"
	fakeAccountEmail    string = "any_valid_email@gmail.com"
	fakeAccountPassword string = "any_password"
)

func NewAccount(userName string, email string, password string) (*Account, error) {
	account := Account{
		Email:    email,
		UserName: userName,
		Password: password,
	}

	return &account, nil
}

func TestNewAccount_WithRightData(t *testing.T) {
	account, err := NewAccount(fakeAccountUserName, fakeAccountEmail, fakeAccountPassword)

	require.Nil(t, err)
	require.Equal(t, account.Email, fakeAccountEmail)
	require.Equal(t, account.UserName, fakeAccountUserName)
	require.Equal(t, account.Password, fakeAccountPassword)
}
