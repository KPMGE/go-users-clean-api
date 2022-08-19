package usecases

import domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"

type GetUserByIdUseCase interface {
	GetUserById(userId string) (*domaindto.GetUserByIdUseCaseOutputDTO, error)
}
