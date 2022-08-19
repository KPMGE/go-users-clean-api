package usecases

import (
	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
)

type AddAccountUseCase interface {
	AddAccount(input *domaindto.AddAccountInputDTO) (*domaindto.AddAccountOutputDTO, error)
}
