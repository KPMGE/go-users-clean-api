package repositories

import (
	"errors"

	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
)

var books []*entities.Book

type InMemoryBookRepository struct{}

func removeIndex[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

func (repo *InMemoryAccountRepository) Add(newBook *entities.Book) (*entities.Book, error) {
	books = append(books, newBook)
	return newBook, nil
}

func (repo *InMemoryAccountRepository) Find(bookId string) (*entities.Book, error) {
	for _, book := range books {
		if book.ID == bookId {
			return book, nil
		}
	}
	return nil, errors.New("book not found!")
}

func (repo *InMemoryAccountRepository) Remove(bookId string) error {
	for index, book := range books {
		if book.ID == book.ID {
			books = removeIndex(books, index)
			return nil
		}
	}
	return errors.New("Cannot find book!")
}

func NewInMemoryBookRepository() *InMemoryAccountRepository {
	return &InMemoryAccountRepository{}
}
