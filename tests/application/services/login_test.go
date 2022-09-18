package usecases_test

import (
	"errors"
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	fakedtos "github.com/KPMGE/go-users-clean-api/tests/domain/fake-dtos"
	"github.com/stretchr/testify/require"
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

func makeLoginServiceSut() (*LoginService, *mocks_test.TokenProviderStub, *mocks_test.FakeAccountRepository) {
	tokenStub := &mocks_test.TokenProviderStub{Output: "some token", Error: nil}
	fakeHasher := mocks_test.NewHasherSpy()
	fakeAccountRepo := mocks_test.NewFakeAccountRepository()
	fakeAccountRepo.CheckAccountOutput = true
	sut := NewLoginService(tokenStub, fakeHasher, fakeAccountRepo)
	return sut, tokenStub, fakeAccountRepo
}

func TestLoginService_ShouldReturnTokenFromProvider(t *testing.T) {
	sut, tokenStub, _ := makeLoginServiceSut()

	token, err := sut.Login(fakedtos.MakeFakeLoginInputDTO())

	require.Nil(t, err)
	require.Equal(t, tokenStub.Output, token)
}

func TestLoginService_ShouldReturnErrorIfTokenProviderReturnsError(t *testing.T) {
	sut, tokenStub, _ := makeLoginServiceSut()
	tokenStub.Error = errors.New("token provider error")

	token, err := sut.Login(fakedtos.MakeFakeLoginInputDTO())

	require.Equal(t, "", token)
	require.Equal(t, tokenStub.Error, err)
}

func TestLoginService_ShouldReturnErrorIfAccountDoesNotExit(t *testing.T) {
	sut, _, accountRepo := makeLoginServiceSut()
	accountRepo.CheckAccountOutput = false

	token, err := sut.Login(fakedtos.MakeFakeLoginInputDTO())

	expectedError := errors.New("account does not exit!")

	require.Equal(t, "", token)
	require.Equal(t, expectedError, err)
}
