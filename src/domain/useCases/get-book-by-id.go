package usecases

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type GetBookByIdUseCase interface {
	GetBookById(bookId string) (*entities.Book, error)
}
