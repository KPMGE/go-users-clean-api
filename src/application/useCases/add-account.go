package usecases

import (
	"errors"
	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	"github.com/KPMGE/go-users-clean-api/src/application/DTO"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
)

type AddAccountUseCase struct {
	accountRepository protocols.AccountRepository
	hasher            protocols.Hasher
}

func (useCase *AddAccountUseCase) AddAccount(input *dto.AddAccountInputDTO) (*dto.AddAccountOutputDTO, error) {
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

	output := dto.AddAccountOutputDTO{
		ID:       account.ID,
		UserName: account.UserName,
		Email:    account.Email,
	}

	return &output, nil
}

func NewAddAccountUseCase(accountRepository protocols.AccountRepository, hasher protocols.Hasher) *AddAccountUseCase {
	return &AddAccountUseCase{
		accountRepository: accountRepository,
		hasher:            hasher,
	}
}
