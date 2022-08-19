package services

import (
	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
)

type ListUsersService struct {
	userRepository protocols.UserRepository
}

func NewListUsersService(repo protocols.UserRepository) *ListUsersService {
	return &ListUsersService{
		userRepository: repo,
	}
}

func (s *ListUsersService) List() []*domaindto.ListUsersDTO {
	users := s.userRepository.List()
	return domaindto.MapListUsersDTO(users)
}
