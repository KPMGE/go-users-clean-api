package usecases

import domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"

type LoginUseCase interface {
	Login(input *domaindto.LoginInputDTO) (string, error)
}
