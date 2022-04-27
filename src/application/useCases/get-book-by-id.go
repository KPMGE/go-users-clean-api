package usecases

import (
	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
)

type GetBookByIdUseCase struct {
	getBookRepo protocols.GetBookRepository
}

func (useCase *GetBookByIdUseCase) GetById(bookId string) (*entities.Book, error) {
	book, err := useCase.getBookRepo.Get(bookId)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func NewGetBookByIdUseCase(getBookRepo protocols.GetBookRepository) *GetBookByIdUseCase {
	return &GetBookByIdUseCase{
		getBookRepo: getBookRepo,
	}
}
