package usecases

import (
	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
)

type ListAccountsUseCase interface {
	ListAccounts() []domaindto.ListAccountsOutputDTO
}
