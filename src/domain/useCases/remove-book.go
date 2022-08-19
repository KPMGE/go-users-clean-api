package usecases

import domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"

type RemoveBookUseCase interface {
	RemoveBook(bookId string) (*domaindto.RemoveBookUseCaseOutputDTO, error)
}
