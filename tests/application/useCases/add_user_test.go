package usecases_test

import (
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

type AddUserOutputDTO struct {
	ID       string
	Name     string
	UserName string
	Email    string
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

	useCase.userRepository.Save(newUser)

	output := AddUserOutputDTO{
		ID:       newUser.ID,
		Name:     newUser.Name,
		UserName: newUser.UserName,
		Email:    newUser.Email,
	}
	return &output, nil
}

func NewAddUserUseCase(repo UserRepository) *AddUserUseCase {
	return &AddUserUseCase{
		userRepository: repo,
	}
}

func makeAddUserSut() (*AddUserUseCase, *UserRepositorySpy) {
	repo := NewUserRepositorySpy()
	sut := NewAddUserUseCase(repo)
	return sut, repo
}

func TestAddUser_WithRightInput(t *testing.T) {
	fakeInput := AddUserInputDTO{
		Name:     "any_name",
		UserName: "any_username",
		Email:    "any_valid_email@gmail.com",
	}

	sut, repo := makeAddUserSut()

	output, err := sut.Add(&fakeInput)

	require.Nil(t, err)
	require.Equal(t, output.Email, fakeInput.Email)
	require.Equal(t, output.UserName, fakeInput.UserName)
	require.Equal(t, output.Name, fakeInput.Name)
	require.NotNil(t, output.ID)
	require.Equal(t, repo.AddInput.Name, fakeInput.Name)
	require.Equal(t, repo.AddInput.Email, fakeInput.Email)
	require.Equal(t, repo.AddInput.UserName, fakeInput.UserName)
}

func TestAddUser_WithWrongInput(t *testing.T) {
	repo := NewUserRepositorySpy()
	sut := NewAddUserUseCase(repo)
	fakeInput := AddUserInputDTO{
		Name:     "any_name",
		UserName: "any_username",
		Email:    "invalid_email",
	}

	output, err := sut.Add(&fakeInput)

	require.Error(t, err)
	require.Equal(t, "Invalid email!", err.Error())
	require.Nil(t, output)
}
