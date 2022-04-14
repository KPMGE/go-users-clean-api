package usecases_test

import (
	"errors"
	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/stretchr/testify/require"
	"testing"
)

type UserRepositorySpy struct {
	AddInput  *entities.User
	AddOutput error
}

func (repo *UserRepositorySpy) Save(user *entities.User) error {
	repo.AddInput = user
	return repo.AddOutput
}

func NewUserRepositorySpy() *UserRepositorySpy {
	return &UserRepositorySpy{}
}

func makeAddUserSut() (*usecases.AddUserUseCase, *UserRepositorySpy) {
	repo := NewUserRepositorySpy()
	repo.AddOutput = nil
	sut := usecases.NewAddUserUseCase(repo)
	return sut, repo
}

func makeFakeValidAddUserInput() *dto.AddUserInputDTO {
	return dto.NewAddUserInputDTO("any_name", "any_username", "any_valid_email@gmail.com")
}

func TestAddUser_WithRightInput(t *testing.T) {
	fakeInput := makeFakeValidAddUserInput()
	sut, repo := makeAddUserSut()

	output, err := sut.Add(fakeInput)

	require.Nil(t, err)
	require.Equal(t, output.Email, fakeInput.Email)
	require.Equal(t, output.UserName, fakeInput.UserName)
	require.Equal(t, output.Name, fakeInput.Name)
	require.NotNil(t, output.ID)
	require.Equal(t, repo.AddInput.Name, fakeInput.Name)
	require.Equal(t, repo.AddInput.Email, fakeInput.Email)
	require.Equal(t, repo.AddInput.UserName, fakeInput.UserName)
}

func TestAddUser_WithInvalidEmail(t *testing.T) {
	repo := NewUserRepositorySpy()
	sut := usecases.NewAddUserUseCase(repo)
	fakeInput := makeFakeValidAddUserInput()
	fakeInput.Email = "invalid_email"

	output, err := sut.Add(fakeInput)

	require.Error(t, err)
	require.Equal(t, "Invalid email!", err.Error())
	require.Nil(t, output)
}

func TestAddUser_WithBlankFields(t *testing.T) {
	repo := NewUserRepositorySpy()
	sut := usecases.NewAddUserUseCase(repo)

	fakeInput := makeFakeValidAddUserInput()
	fakeInput.Name = ""
	output, err := sut.Add(fakeInput)
	require.Error(t, err)
	require.Nil(t, output)

	fakeInput = makeFakeValidAddUserInput()
	fakeInput.UserName = ""
	output, err = sut.Add(fakeInput)
	require.Error(t, err)
	require.Nil(t, output)

	fakeInput = makeFakeValidAddUserInput()
	fakeInput.Email = ""
	output, err = sut.Add(fakeInput)
	require.Error(t, err)
	require.Nil(t, output)
}

func TestAddUser_WhenRepositoryReturnsError(t *testing.T) {
	repo := NewUserRepositorySpy()
	repo.AddOutput = errors.New("some error")
	sut := usecases.NewAddUserUseCase(repo)
	fakeInput := makeFakeValidAddUserInput()

	output, err := sut.Add(fakeInput)

	require.Nil(t, output)
	require.Error(t, err)
	require.Equal(t, err.Error(), "some error")
}
