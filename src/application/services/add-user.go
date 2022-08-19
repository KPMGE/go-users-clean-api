package services

import (
	"errors"

	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
)

type AddUserService struct {
	userRepository protocols.UserRepository
}

func (s *AddUserService) Add(input *domaindto.AddUserInputDTO) (*domaindto.AddUserOutputDTO, error) {
	newUser, err := entities.NewUser(input.Name, input.UserName, input.Email)
	if err != nil {
		return nil, err
	}

	emailTaken := s.userRepository.CheckByEmail(input.Email)
	if emailTaken {
		return nil, errors.New("email already taken!")
	}

	userNameTaken := s.userRepository.CheckByUserName(input.UserName)
	if userNameTaken {
		return nil, errors.New("UserName already taken!")
	}

	err = s.userRepository.Save(newUser)
	if err != nil {
		return nil, err
	}

	output := domaindto.NewAddUserOutputDTO(newUser.ID, newUser.Name, newUser.UserName, newUser.Email)
	return output, nil
}

func NewAddUserService(repo protocols.UserRepository) *AddUserService {
	return &AddUserService{
		userRepository: repo,
	}
}
