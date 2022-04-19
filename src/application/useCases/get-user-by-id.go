package usecases

import (
	"errors"

	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
)

type GetUserByIdUseCase struct {
	userRepository protocols.UserRepository
}

func NewGetUserByIdUseCase(repo protocols.UserRepository) *GetUserByIdUseCase {
	return &GetUserByIdUseCase{
		userRepository: repo,
	}
}

func (useCase *GetUserByIdUseCase) Get(userId string) (*dto.GetUserByIdUseCaseOutputDTO, error) {
	foundUser, err := useCase.userRepository.GetById(userId)

	if err != nil {
		return nil, err
	}

	if foundUser == nil {
		return nil, errors.New("User not found!")
	}

	output := dto.NewGetUserByIdUseCaseOutputDTO(foundUser.ID, foundUser.Name, foundUser.Email, foundUser.UserName)
	return output, nil
}
