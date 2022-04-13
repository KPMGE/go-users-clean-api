package usecases_test

import (
	"testing"

	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

const fakeUserName string = "any_user_name"
const fakeEmail string = "any_valid_email@gmail.com"
const fakePassword string = "any_password"

func makeFakeInput() *dto.AddAccountInputDTO {
	return &dto.AddAccountInputDTO{
		UserName:        fakeUserName,
		Email:           fakeEmail,
		Password:        fakePassword,
		ConfirmPassword: fakePassword,
	}
}

func makeSut() (*usecases.AddAccountUseCase, *mocks_test.HasherSpy, *mocks_test.FakeAccountRepository) {
	repo := mocks_test.NewFakeAccountRepository()
	hasher := mocks_test.NewHasherSpy()
	sut := usecases.NewAddAccountUseCase(repo, hasher)
	return sut, hasher, repo
}

func TestAddAccountUseCase_WithRightData(t *testing.T) {
	sut, hasher, repo := makeSut()
	fakeInput := makeFakeInput()
	createdAccount, err := sut.AddAccount(fakeInput)

	require.Nil(t, err)
	require.Equal(t, hasher.Input, fakePassword)
	require.Equal(t, createdAccount.Email, fakeEmail)
	require.Equal(t, createdAccount.UserName, fakeUserName)
	require.Equal(t, repo.Input.Password, "hashed_text")
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
	repo.CheckAccountOutput = true
	fakeInput := makeFakeInput()

	createdAccount, err := sut.AddAccount(fakeInput)

	require.Error(t, err)
	require.Nil(t, createdAccount)
	require.Equal(t, err.Error(), "email already taken")
}

func TestAddAccountUseCase_WithUsernameAlreadyTaken(t *testing.T) {
	sut, _, repo := makeSut()
	fakeInput := makeFakeInput()
	repo.CheckUserNameOutput = true

	createdAccount, err := sut.AddAccount(fakeInput)

	require.Error(t, err)
	require.Nil(t, createdAccount)
	require.Equal(t, err.Error(), "username already taken")
}

func TestAddAccountUseCase_WithBlankFields(t *testing.T) {
	sut, _, repo := makeSut()
	repo.CheckUserNameOutput = false
	repo.CheckUserNameOutput = false

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
