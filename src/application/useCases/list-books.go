package usecases

import (
	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
)

type ListBooksUseCase struct {
	listBooksRepo protocols.ListBooksRepository
}

func NewListBookUseCase(repo protocols.ListBooksRepository) *ListBooksUseCase {
	return &ListBooksUseCase{
		listBooksRepo: repo,
	}
}

func (useCase *ListBooksUseCase) List() ([]*entities.Book, error) {
	books, err := useCase.listBooksRepo.List()
	if err != nil {
		return nil, err
	}
	return books, nil
}
