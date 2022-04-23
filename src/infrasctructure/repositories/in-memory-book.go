package repositories

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

var books []*entities.Book

type InMemoryBookRepository struct{}

func (repo *InMemoryAccountRepository) Add(newBook *entities.Book) (*entities.Book, error) {
	books = append(books, newBook)
	return newBook, nil
}

func NewInMemoryBookRepository() *InMemoryAccountRepository {
	return &InMemoryAccountRepository{}
}
