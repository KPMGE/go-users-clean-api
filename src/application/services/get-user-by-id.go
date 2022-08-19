package services

import (
	"errors"

	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
)

type GetUserByIdService struct {
	userRepository protocols.UserRepository
}

func NewGetUserByIdService(repo protocols.UserRepository) *GetUserByIdService {
	return &GetUserByIdService{
		userRepository: repo,
	}
}

func (s *GetUserByIdService) GetUserById(userId string) (*domaindto.GetUserByIdUseCaseOutputDTO, error) {
	foundUser, err := s.userRepository.GetById(userId)

	if err != nil {
		return nil, err
	}

	if foundUser == nil {
		return nil, errors.New("User not found!")
	}

	output := domaindto.NewGetUserByIdUseCaseOutputDTO(
		foundUser.ID,
		foundUser.Name,
		foundUser.Email,
		foundUser.UserName,
		foundUser.Books,
	)

	return output, nil
}
