package services

import (
	"errors"

	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
)

type AddAccountService struct {
	accountRepository protocols.AccountRepository
	hasher            protocols.Hasher
}

func (useCase *AddAccountService) AddAccount(input *domaindto.AddAccountInputDTO) (*domaindto.AddAccountOutputDTO, error) {
	emailTaken := useCase.accountRepository.CheckAccountByEmail(input.Email)
	if emailTaken {
		return nil, errors.New("email already taken")
	}

	userNameTaken := useCase.accountRepository.CheckAccountByUserName(input.UserName)
	if userNameTaken {
		return nil, errors.New("username already taken")
	}

	if input.Password != input.ConfirmPassword {
		return nil, errors.New("password and confirmPassword must match")
	}

	hashedPassword := useCase.hasher.Hash(input.Password)
	account, err := entities.NewAccount(input.UserName, input.Email, hashedPassword)

	if err != nil {
		return nil, err
	}

	err = useCase.accountRepository.Save(account)
	if err != nil {
		return nil, err
	}

	output := domaindto.NewAddAccountOutputDTO(account.ID, account.UserName, account.Email)

	return output, nil
}

func NewAddAccountService(accountRepository protocols.AccountRepository, hasher protocols.Hasher) *AddAccountService {
	return &AddAccountService{
		accountRepository: accountRepository,
		hasher:            hasher,
	}
}
