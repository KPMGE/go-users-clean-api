package usecases

import domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"

type AddUserUseCase interface {
	Add(input *domaindto.AddUserInputDTO) (*domaindto.AddUserOutputDTO, error)
}
