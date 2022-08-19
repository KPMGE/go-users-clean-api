package usecases

import domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"

type ListUsersUseCase interface {
	List() []*domaindto.ListUsersDTO
}
