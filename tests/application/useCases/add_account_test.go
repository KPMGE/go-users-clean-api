package usecases_test

import (
	"testing"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/stretchr/testify/require"
)

type hasherSpy struct {
	input string
}

func (hasher *hasherSpy) Hash(plainText string) string {
	hasher.input = plainText
	return "hashed_text"
}

type FakeAccountRepository struct {
	input               *entities.Account
	checkAccountOutput  bool
	checkUserNameOutput bool
}

func (repo *FakeAccountRepository) CheckAccountByEmail(email string) bool {
	return repo.checkAccountOutput
}

func (repo *FakeAccountRepository) CheckAccountByUserName(userName string) bool {
	return repo.checkUserNameOutput
}

func (repo *FakeAccountRepository) Save(account *entities.Account) error {
	repo.input = account
	return nil
}

func NewHasherSpy() *hasherSpy {
	return &hasherSpy{}
}

func NewFakeAccountRepository() *FakeAccountRepository {
	return &FakeAccountRepository{}
}

const fakeUserName string = "any_user_name"
const fakeEmail string = "any_valid_email@gmail.com"
const fakePassword string = "any_password"

func makeFakeInput() *usecases.AddAccountInputDTO {
	return &usecases.AddAccountInputDTO{
		UserName:        fakeUserName,
		Email:           fakeEmail,
		Password:        fakePassword,
		ConfirmPassword: fakePassword,
	}
}

func makeSut() (*usecases.AddAccountUseCase, *hasherSpy, *FakeAccountRepository) {
	repo := NewFakeAccountRepository()
	hasher := NewHasherSpy()
	sut := usecases.NewAddAccountUseCase(repo, hasher)
	return sut, hasher, repo
}

func TestAddAccountUseCase_WithRightData(t *testing.T) {
	sut, hasher, repo := makeSut()
	fakeInput := makeFakeInput()
	createdAccount, err := sut.AddAccount(fakeInput)

	require.Nil(t, err)
	require.Equal(t, hasher.input, fakePassword)
	require.Equal(t, createdAccount.Email, fakeEmail)
	require.Equal(t, createdAccount.UserName, fakeUserName)
	require.Equal(t, repo.input.Password, "hashed_text")
}

func TestAddAccountUseCase_WithDifferentPasswordAndConfirmPassword(t *testing.T) {
	sut, _, _ := makeSut()
	fakeInput := makeFakeInput()
	fakeInput.ConfirmPassword = "any_different_password"

	createdAccount, err := sut.AddAccount(fakeInput)

	require.Error(t, err)
	require.Nil(t, createdAccount)
	require.Equal(t, err.Error(), "password and confirmPassword must match")
}

func TestAddAccountUseCase_WithEmailAlreadyTaken(t *testing.T) {
	sut, _, repo := makeSut()
	repo.checkAccountOutput = true
	fakeInput := makeFakeInput()

	createdAccount, err := sut.AddAccount(fakeInput)

	require.Error(t, err)
	require.Nil(t, createdAccount)
	require.Equal(t, err.Error(), "email already taken")
}

func TestAddAccountUseCase_WithUsernameAlreadyTaken(t *testing.T) {
	sut, _, repo := makeSut()
	fakeInput := makeFakeInput()
	repo.checkUserNameOutput = true

	createdAccount, err := sut.AddAccount(fakeInput)

	require.Error(t, err)
	require.Nil(t, createdAccount)
	require.Equal(t, err.Error(), "username already taken")
}

func TestAddAccountUseCase_WithBlankFields(t *testing.T) {
	sut, _, repo := makeSut()
	repo.checkUserNameOutput = false
	repo.checkUserNameOutput = false

	fakeInput := makeFakeInput()
	fakeInput.UserName = ""
	createdAccount, err := sut.AddAccount(fakeInput)
	require.Error(t, err)
	require.Nil(t, createdAccount)

	fakeInput = makeFakeInput()
	fakeInput.Password = ""
	createdAccount, err = sut.AddAccount(fakeInput)
	require.Error(t, err)
	require.Nil(t, createdAccount)

	fakeInput = makeFakeInput()
	fakeInput.ConfirmPassword = ""
	createdAccount, err = sut.AddAccount(fakeInput)
	require.Error(t, err)
	require.Nil(t, createdAccount)

	fakeInput = makeFakeInput()
	fakeInput.Email = ""
	createdAccount, err = sut.AddAccount(fakeInput)
	require.Error(t, err)
	require.Nil(t, createdAccount)
}
