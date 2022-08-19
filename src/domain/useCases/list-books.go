package usecases

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type ListBooksUseCase interface {
	ListBooks() ([]*entities.Book, error)
}
