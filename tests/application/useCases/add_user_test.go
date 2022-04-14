package usecases_test

import (
	"errors"
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/stretchr/testify/require"
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

type AddUserInputDTO struct {
	Name     string
	UserName string
	Email    string
}

func NewAddUserInputDTO(name string, userName string, email string) *AddUserInputDTO {
	return &AddUserInputDTO{
		Name:     name,
		Email:    email,
		UserName: userName,
	}
}

type AddUserOutputDTO struct {
	ID       string
	Name     string
	UserName string
	Email    string
}

func NewAddUserOutputDTO(id string, name string, userName string, email string) *AddUserOutputDTO {
	return &AddUserOutputDTO{
		Name:     name,
		Email:    email,
		ID:       id,
		UserName: userName,
	}
}

type UserRepository interface {
	Save(user *entities.User) error
}

type AddUserUseCase struct {
	userRepository UserRepository
}

func (useCase *AddUserUseCase) Add(input *AddUserInputDTO) (*AddUserOutputDTO, error) {
	newUser, err := entities.NewUser(input.Name, input.UserName, input.Email)
	if err != nil {
		return nil, err
	}

	err = useCase.userRepository.Save(newUser)
	if err != nil {
		return nil, err
	}

	output := NewAddUserOutputDTO(newUser.ID, newUser.Name, newUser.UserName, newUser.Email)
	return output, nil
}

func NewAddUserUseCase(repo UserRepository) *AddUserUseCase {
	return &AddUserUseCase{
		userRepository: repo,
	}
}

func makeAddUserSut() (*AddUserUseCase, *UserRepositorySpy) {
	repo := NewUserRepositorySpy()
	repo.AddOutput = nil
	sut := NewAddUserUseCase(repo)
	return sut, repo
}

func makeFakeValidAddUserInput() *AddUserInputDTO {
	return NewAddUserInputDTO("any_name", "any_username", "any_valid_email@gmail.com")
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
	sut := NewAddUserUseCase(repo)
	fakeInput := makeFakeValidAddUserInput()
	fakeInput.Email = "invalid_email"

	output, err := sut.Add(fakeInput)

	require.Error(t, err)
	require.Equal(t, "Invalid email!", err.Error())
	require.Nil(t, output)
}

func TestAddUser_WithBlankFields(t *testing.T) {
	repo := NewUserRepositorySpy()
	sut := NewAddUserUseCase(repo)

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
	sut := NewAddUserUseCase(repo)
	fakeInput := makeFakeValidAddUserInput()

	output, err := sut.Add(fakeInput)

	require.Nil(t, output)
	require.Error(t, err)
	require.Equal(t, err.Error(), "some error")
}
