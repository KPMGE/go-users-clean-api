package usecases_test

import (
	"errors"
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

type TokenProvider interface {
	generate(data any) (string, error)
}

type TokenProviderStub struct {
	Output string
	Error  error
}

type LoginInputDTO struct {
	UserName string
	Email    string
	Password string
}

func makeFakeLoginInputDTO() *LoginInputDTO {
	return &LoginInputDTO{
		Email:    "any@email.com",
		Password: "any password",
	}
}

type LoginUseCase interface {
	Login(input *LoginInputDTO) (string, error)
}

type LoginService struct {
	tk   TokenProvider
	hs   protocols.Hasher
	repo protocols.AccountRepository
}

func (l *LoginService) Login(input *LoginInputDTO) (string, error) {
	accountExists := l.repo.CheckAccountByEmail(input.Email)

	if !accountExists {
		return "", errors.New("account does not exit!")
	}

	hashedPassword := l.hs.Hash(input.Password)
	input.Password = hashedPassword

	token, err := l.tk.generate(input)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (t *TokenProviderStub) generate(data any) (string, error) {
	return t.Output, t.Error
}

func NewLoginService(tk TokenProvider, hs protocols.Hasher, repo protocols.AccountRepository) *LoginService {
	return &LoginService{
		tk:   tk,
		hs:   hs,
		repo: repo,
	}
}

func makeLoginServiceSut() (*LoginService, *TokenProviderStub, *mocks_test.FakeAccountRepository) {
	tokenStub := &TokenProviderStub{Output: "some token", Error: nil}
	fakeHasher := mocks_test.NewHasherSpy()
	fakeAccountRepo := mocks_test.NewFakeAccountRepository()
	fakeAccountRepo.CheckAccountOutput = true
	sut := NewLoginService(tokenStub, fakeHasher, fakeAccountRepo)
	return sut, tokenStub, fakeAccountRepo
}

func TestLoginService_ShouldReturnTokenFromProvider(t *testing.T) {
	sut, tokenStub, _ := makeLoginServiceSut()

	token, err := sut.Login(makeFakeLoginInputDTO())

	require.Nil(t, err)
	require.Equal(t, tokenStub.Output, token)
}

func TestLoginService_ShouldReturnErrorIfTokenProviderReturnsError(t *testing.T) {
	sut, tokenStub, _ := makeLoginServiceSut()
	tokenStub.Error = errors.New("token provider error")

	token, err := sut.Login(makeFakeLoginInputDTO())

	require.Equal(t, "", token)
	require.Equal(t, tokenStub.Error, err)
}

func TestLoginService_ShouldReturnErrorIfAccountDoesNotExit(t *testing.T) {
	sut, _, accountRepo := makeLoginServiceSut()
	accountRepo.CheckAccountOutput = false

	token, err := sut.Login(makeFakeLoginInputDTO())

	expectedError := errors.New("account does not exit!")

	require.Equal(t, "", token)
	require.Equal(t, expectedError, err)
}
