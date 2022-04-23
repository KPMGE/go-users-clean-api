package entities

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Book struct {
	Base
	Title       string  `json:"title" valid:"required"`
	Author      string  `json:"author" valid:"required"`
	Price       float64 `json:"price" valid:"required"`
	Description string  `json:"description" valid:"required"`
	UserId      string  `json:"userId" valid:"required"`
}

func (book *Book) isValid() error {
	if book.Price <= 0 {
		return errors.New("Price must be greater than 0!")
	}
	_, err := govalidator.ValidateStruct(book)
	if err != nil {
		return err
	}
	return nil
}

func NewBook(title string, author string, description string, price float64, UserId string) (*Book, error) {
	book := Book{
		UserId:      UserId,
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
