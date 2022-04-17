package usecases_test

import (
	"errors"
	"log"
	"testing"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

const fakeName string = "any_user_name"

func MakeGetUserByIdSut() (*usecases.GetUserByIdUseCase, *mocks_test.UserRepositorySpy) {
	repo := mocks_test.NewUserRepositorySpy()
	fakeUser, err := entities.NewUser(fakeName, fakeUserName, fakeEmail)
	if err != nil {
		log.Fatal(err)
	}
	repo.GetByidOutput = fakeUser
	sut := usecases.NewGetUserByIdUseCase(repo)
	return sut, repo
}

func TestGetUserByIdUseCase_ShouldCallRepositoryCorrectly(t *testing.T) {
	sut, repo := MakeGetUserByIdSut()

	sut.Get(FAKE_USER_ID)

	require.Equal(t, FAKE_USER_ID, repo.GetByidInput)
}

func TestGetUserByIdUseCase_ShouldReturnFoundUser(t *testing.T) {
	sut, _ := MakeGetUserByIdSut()

	foundUser, _ := sut.Get(FAKE_USER_ID)

	require.Equal(t, fakeName, foundUser.Name)
	require.Equal(t, fakeUserName, foundUser.UserName)
	require.Equal(t, fakeEmail, foundUser.Email)
	require.NotNil(t, foundUser.ID)
}

func TestGetUserByIdUseCase_ShouldReturnErrorIfRepositoryRetunsNil(t *testing.T) {
	sut, repo := MakeGetUserByIdSut()
	repo.GetByidOutput = nil

	foundUser, err := sut.Get("invalid_id")

	require.Error(t, err)
	require.Equal(t, "User not found!", err.Error())
	require.Nil(t, foundUser)
}

func TestGetUserByIdUseCase_ShouldReturnErrorIfRepositoryRetunsError(t *testing.T) {
	sut, repo := MakeGetUserByIdSut()
	repo.GetByidError = errors.New("repository error")

	foundUser, err := sut.Get("invalid_id")

	require.Error(t, err)
	require.Equal(t, "repository error", err.Error())
	require.Nil(t, foundUser)
}
