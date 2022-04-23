package entities_test

import (
	"testing"
	"time"

	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

const (
	fakeTitle       string  = "any_title"
	fakeAuthor      string  = "any_author"
	fakeDescription string  = "any_description"
	fakePrice       float64 = 5.3
)

func makeFakeUser() *entities.User {
	user := entities.User{
		Name:     "any_name",
		UserName: "any_user_name",
		Email:    "any_valid_email@gmail.com",
	}

	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return &user
}

func TestNewBook_WithRighData(t *testing.T) {
	fakeUser := makeFakeUser()
	newBook, err := entities.NewBook(fakeTitle, fakeAuthor, fakeDescription, fakePrice, fakeUser.ID)

	require.Nil(t, err)
	require.NotEmpty(t, newBook.CreatedAt)
	require.NotEmpty(t, newBook.UpdatedAt)
	require.NotEmpty(t, newBook.ID)
	require.Equal(t, newBook.Title, fakeTitle)
	require.Equal(t, newBook.Author, fakeAuthor)
	require.Equal(t, newBook.Price, fakePrice)
	require.Equal(t, newBook.Description, fakeDescription)
}

func TestNewBook_WithPriceLessThanOrEqualTo0(t *testing.T) {
	fakeUser := makeFakeUser()
	newBook, err := entities.NewBook(fakeTitle, fakeAuthor, fakeDescription, 0, fakeUser.ID)

	require.Error(t, err)
	require.Nil(t, newBook)
	require.Equal(t, err.Error(), "Price must be greater than 0!")
}

func TestNewBook_WithNullFields(t *testing.T) {
	fakeUser := makeFakeUser()

	newBook, err := entities.NewBook("", fakeAuthor, fakeDescription, fakePrice, fakeUser.ID)
	require.Error(t, err)
	require.Nil(t, newBook)

	newBook, err = entities.NewBook(fakeTitle, "", fakeDescription, fakePrice, fakeUser.ID)
	require.Error(t, err)
	require.Nil(t, newBook)

	newBook, err = entities.NewBook(fakeTitle, fakeAuthor, "", fakePrice, fakeUser.ID)
	require.Error(t, err)
	require.Nil(t, newBook)
}
