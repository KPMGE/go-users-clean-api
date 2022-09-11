package services

import (
	"errors"
	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
)

type DeleteUserService struct {
	userRepository protocols.UserRepository
}

func (s *DeleteUserService) DeleteUser(userId string) (string, error) {
	userExists := s.userRepository.CheckById(userId)
	if !userExists {
		return "", errors.New("there is no user with the provided id")
	}
	err := s.userRepository.Delete(userId)
	if err != nil {
		return "", err
	}
	return "user deleted successfully", nil
}

func NewDeleteUserService(repo protocols.UserRepository) *DeleteUserService {
	return &DeleteUserService{
		userRepository: repo,
	}
}
