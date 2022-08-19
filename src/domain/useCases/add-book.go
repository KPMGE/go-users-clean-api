package usecases

import domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"

type AddBookUseCase interface {
	AddBook(input *domaindto.AddBookUseCaseInputDTO) (*domaindto.AddBookUseCaseOutputDTO, error)
}
