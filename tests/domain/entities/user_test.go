package entities_test

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type Base struct {
	ID        string    `json:"id" valid:"uuid"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

type User struct {
	Base
	Name     string  `json:"name" valid:"required"`
	UserName string  `json:"user_name" valid:"required"`
	Email    string  `json:"email" valid:"email"`
	Books    []*Book `json:"books" valid:"-"`
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
	user.Books = nil
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err := user.isValid()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

const fakeName string = "any_name"
const fakeUserName string = "any_user_name"
const fakeEmail string = "any_valid_email@gmail.com"

func TestUser_New(t *testing.T) {
	newUser, err := NewUser(fakeName, fakeUserName, fakeEmail)

	require.Nil(t, err)
	require.Nil(t, newUser.Books)
	require.NotNil(t, newUser)
	require.NotNil(t, newUser.ID)
	require.NotNil(t, newUser.UpdatedAt)
	require.NotNil(t, newUser.CreatedAt)
	require.Equal(t, fakeName, newUser.Name)
	require.Equal(t, fakeUserName, newUser.UserName)
	require.Equal(t, fakeEmail, newUser.Email)
}

func TestUser_New_WithInvalidEmail(t *testing.T) {
	newUser, err := NewUser(fakeName, fakeUserName, "any_invalid_email")

	require.Error(t, err)
	require.Nil(t, newUser)
	require.Equal(t, err.Error(), "Invalid email!")
}

func TestUser_New_WithNullFields(t *testing.T) {
	newUser, err := NewUser("", fakeUserName, fakeEmail)
	require.Error(t, err)
	require.Nil(t, newUser)

	newUser, err = NewUser(fakeName, "", fakeEmail)
	require.Error(t, err)
	require.Nil(t, newUser)

	newUser, err = NewUser(fakeName, fakeUserName, "")
	require.Error(t, err)
	require.Nil(t, newUser)
}
