package entities_test

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/stretchr/testify/require"
	"testing"
)

type Account struct {
	UserName string `json:"user_name" valid:"required"`
	Email    string `json:"email" valid:"required"`
	Password string `json:"password" valid:"required"`
}

const (
	fakeAccountUserName string = "any_name"
	fakeAccountEmail    string = "any_valid_email@gmail.com"
	fakeAccountPassword string = "any_password"
)

func (account *Account) isValid() error {
	validEmail := govalidator.IsEmail(account.Email)
	if !validEmail {
		return errors.New("Invalid email!")
	}

	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}

	return nil
}

func NewAccount(userName string, email string, password string) (*Account, error) {
	account := Account{
		Email:    email,
		UserName: userName,
		Password: password,
	}

	err := account.isValid()
	if err != nil {
		return nil, err
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

func TestNewAccount_WithInvalidEmail(t *testing.T) {
	account, err := NewAccount(fakeAccountUserName, "any_invalid_email", fakeAccountPassword)

	require.Nil(t, account)
	require.Error(t, err)
	require.Equal(t, err.Error(), "Invalid email!")
}

func TestNewAccount_WithBlankFields(t *testing.T) {
	account, err := NewAccount("", fakeAccountEmail, fakeAccountPassword)
	require.Nil(t, account)
	require.Error(t, err)

	account, err = NewAccount(fakeAccountUserName, "", fakeAccountPassword)
	require.Nil(t, account)
	require.Error(t, err)

	account, err = NewAccount(fakeAccountUserName, fakeAccountEmail, "")
	require.Nil(t, account)
	require.Error(t, err)
}
