package entities_test

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type Book struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Price     float64   `json:"price"`
	User      *User     `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	Books     []*Book   `json:"books"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

	return &user, nil
}

func TestUser_New(t *testing.T) {
	fakeName := "any_name"
	fakeUserName := "any_user_name"
	fakeEmail := "any_valid_email@gmail.com"

	newUser, err := NewUser(fakeName, fakeUserName, fakeEmail)

	require.Nil(t, err)
	require.Nil(t, newUser.Books)
	require.NotNil(t, newUser)
	require.NotNil(t, newUser.ID)
	require.NotNil(t, newUser.UpdatedAt)
	require.NotNil(t, newUser.CreatedAt)
}
