package usecases

import (
	"errors"

	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
)

type AddUserUseCase struct {
	userRepository protocols.UserRepository
}

func (useCase *AddUserUseCase) Add(input *dto.AddUserInputDTO) (*dto.AddUserOutputDTO, error) {
	newUser, err := entities.NewUser(input.Name, input.UserName, input.Email)
	if err != nil {
		return nil, err
	}

	emailTaken := useCase.userRepository.CheckByEmail(input.Email)
	if emailTaken {
		return nil, errors.New("email already taken!")
	}

	userNameTaken := useCase.userRepository.CheckByUserName(input.UserName)
	if userNameTaken {
		return nil, errors.New("UserName already taken!")
	}

	err = useCase.userRepository.Save(newUser)
	if err != nil {
		return nil, err
	}

	output := dto.NewAddUserOutputDTO(newUser.ID, newUser.Name, newUser.UserName, newUser.Email)
	return output, nil
}

func NewAddUserUseCase(repo protocols.UserRepository) *AddUserUseCase {
	return &AddUserUseCase{
		userRepository: repo,
	}
}
