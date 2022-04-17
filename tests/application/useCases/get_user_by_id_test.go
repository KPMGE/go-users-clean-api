package usecases_test

import (
	"log"
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

const fakeName string = "any_user_name"

type GetUserByIdUseCaseOutputDTO struct {
	ID       string
	Name     string
	Email    string
	UserName string
}

type GetUserByIdUseCase struct {
	userRepository protocols.UserRepository
}

func NewGetUserByIdUseCase(repo protocols.UserRepository) *GetUserByIdUseCase {
	return &GetUserByIdUseCase{
		userRepository: repo,
	}
}

func (useCase *GetUserByIdUseCase) Get(userId string) *GetUserByIdUseCaseOutputDTO {
	foundUser, _ := useCase.userRepository.GetById(userId)
	return &GetUserByIdUseCaseOutputDTO{
		ID:       foundUser.ID,
		Name:     foundUser.Name,
		UserName: foundUser.UserName,
		Email:    foundUser.Email,
	}
}

func MakeGetUserByIdSut() (*GetUserByIdUseCase, *mocks_test.UserRepositorySpy) {
	repo := mocks_test.NewUserRepositorySpy()
	fakeUser, err := entities.NewUser(fakeName, fakeUserName, fakeEmail)
	if err != nil {
		log.Fatal(err)
	}
	repo.GetByidOutput = fakeUser
	sut := NewGetUserByIdUseCase(repo)
	return sut, repo
}

func TestGetUserByIdUseCase_ShouldCallRepositoryCorrectly(t *testing.T) {
	sut, repo := MakeGetUserByIdSut()
	sut.Get(FAKE_USER_ID)
	require.Equal(t, FAKE_USER_ID, repo.GetByidInput)
}

func TestGetUserByIdUseCase_ShouldReturnFoundUser(t *testing.T) {
	sut, _ := MakeGetUserByIdSut()
	foundUser := sut.Get(FAKE_USER_ID)
	require.Equal(t, fakeName, foundUser.Name)
	require.Equal(t, fakeUserName, foundUser.UserName)
	require.Equal(t, fakeEmail, foundUser.Email)
	require.NotNil(t, foundUser.ID)
}
