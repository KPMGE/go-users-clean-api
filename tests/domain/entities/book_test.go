package entities_test

import (
	"errors"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

type Book struct {
	Base
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	User        *User   `json:"user"`
}

const (
	fakeTitle       string  = "any_title"
	fakeAuthor      string  = "any_author"
	fakeDescription string  = "any_description"
	fakePrice       float64 = 5.3
)

func makeFakeUser() *User {
	user := User{
		Name:     "any_name",
		UserName: "any_user_name",
		Email:    "any_valid_email@gmail.com",
	}

	user.ID = "any_id"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return &user
}

func (book *Book) isValid() error {
	if book.Price <= 0 {
		return errors.New("Price must be greater than 0!")
	}
	return nil
}

func NewBook(title string, author string, description string, price float64, user *User) (*Book, error) {
	book := Book{
		User:        user,
		Title:       title,
		Author:      author,
		Description: description,
		Price:       price,
	}

	book.ID = uuid.NewV4().String()
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	err := book.isValid()
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func TestNewBook_WithRighData(t *testing.T) {
	fakeUser := makeFakeUser()
	newBook, err := NewBook(fakeTitle, fakeAuthor, fakeDescription, fakePrice, fakeUser)

	require.Nil(t, err)
	require.NotEmpty(t, newBook.CreatedAt)
	require.NotEmpty(t, newBook.UpdatedAt)
	require.NotEmpty(t, newBook.ID)
	require.Equal(t, newBook.User, fakeUser)
	require.Equal(t, newBook.Title, fakeTitle)
	require.Equal(t, newBook.Author, fakeAuthor)
	require.Equal(t, newBook.Price, fakePrice)
	require.Equal(t, newBook.Description, fakeDescription)
}

func TestNewBook_WithPriceLessThanOrEqualTo0(t *testing.T) {
	fakeUser := makeFakeUser()
	newBook, err := NewBook(fakeTitle, fakeAuthor, fakeDescription, 0, fakeUser)

	require.Error(t, err)
	require.Nil(t, newBook)
	require.Equal(t, err.Error(), "Price must be greater than 0!")
}
