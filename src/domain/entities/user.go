package entities

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	Base
	Name     string `json:"name" valid:"required"`
	UserName string `json:"user_name" valid:"required"`
	Email    string `json:"email" valid:"email"`
	Books    []Book `json:"books" valid:"-"`
}

func (user *User) isValid() error {
	validEmail := govalidator.IsEmail(user.Email)
	if !validEmail {
		return errors.New("Invalid email!")
	}

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}

	return nil
}

func NewUser(name string, userName string, email string) (*User, error) {
	user := User{
		Name:     name,
		UserName: userName,
		Email:    email,
	}

	user.ID = uuid.NewV4().String()
	user.Books = []Book{}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err := user.isValid()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
