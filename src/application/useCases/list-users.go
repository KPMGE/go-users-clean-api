package usecases

import (
	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
)

type ListUsersUseCase struct {
	userRepository protocols.UserRepository
}

func NewListUsersUseCase(repo protocols.UserRepository) *ListUsersUseCase {
	return &ListUsersUseCase{
		userRepository: repo,
	}
}

func (useCase *ListUsersUseCase) List() []*dto.ListUsersDTO {
	users := useCase.userRepository.List()
	return dto.MapListUsersDTO(users)
}
