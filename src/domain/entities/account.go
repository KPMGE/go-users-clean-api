package entities

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Base
	UserName string `json:"user_name" valid:"required"`
	Email    string `json:"email" valid:"required"`
	Password string `json:"password" valid:"required"`
}

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

	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()
	account.ID = uuid.NewV4().String()

	err := account.isValid()
	if err != nil {
		return nil, err
	}

	return &account, nil
}
