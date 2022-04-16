package usecases

import (
	"errors"

	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
)

type DeleteUserUseCase struct {
	userRepository protocols.UserRepository
}

func NewDeleteUserUseCase(repo protocols.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		userRepository: repo,
	}
}

func (useCase *DeleteUserUseCase) Delete(userId string) (string, error) {
	userExists := useCase.userRepository.CheckById(userId)
	if !userExists {
		return "", errors.New("No user with the provided id!")
	}
	err := useCase.userRepository.Delete(userId)
	if err != nil {
		return "", err
	}
	return "user deleted successfully", nil
}
