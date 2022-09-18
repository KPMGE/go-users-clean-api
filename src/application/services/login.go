package services

import (
	"errors"

	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
)

type LoginService struct {
	tk   protocols.TokenProvider
	hs   protocols.Hasher
	repo protocols.AccountRepository
}

func (l *LoginService) Login(input *domaindto.LoginInputDTO) (string, error) {
	accountExists := l.repo.CheckAccountByEmail(input.Email)

	if !accountExists {
		return "", errors.New("account does not exit!")
	}

	hashedPassword := l.hs.Hash(input.Password)
	input.Password = hashedPassword

	token, err := l.tk.Generate(input)

	if err != nil {
		return "", err
	}

	return token, nil
}

func NewLoginService(tk protocols.TokenProvider, hs protocols.Hasher, repo protocols.AccountRepository) *LoginService {
	return &LoginService{
		tk:   tk,
		hs:   hs,
		repo: repo,
	}
}
