package usecases_test

import (
	"errors"
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

const FAKE_USER_ID string = "any_user_id"

type DeleteUserUseCase struct {
	userRepository protocols.UserRepository
}

func NewDeleteUserUseCase(repo protocols.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		userRepository: repo,
	}
}

func (useCase *DeleteUserUseCase) Delete(userId string) (string, error) {
	userExists := useCase.userRepository.CheckById(userId)
	if !userExists {
		return "", errors.New("No user with the provided id!")
	}
	err := useCase.userRepository.Delete(userId)
	if err != nil {
		return "", err
	}
	return "user deleted successfully", nil
}

func MakeDeleteUserSut() (*DeleteUserUseCase, *mocks_test.UserRepositorySpy) {
	repo := mocks_test.NewUserRepositorySpy()
	repo.DeleteOutput = nil
	repo.CheckByIdOuput = true
	sut := NewDeleteUserUseCase(repo)
	return sut, repo
}

func TestDeleteUserUseCase_ShouldReturnRightDataOnSuccess(t *testing.T) {
	sut, _ := MakeDeleteUserSut()
	message, err := sut.Delete(FAKE_USER_ID)

	require.Nil(t, err)
	require.Equal(t, "user deleted successfully", message)
}

func TestDeleteUserUseCase_ShouldCallUserChechUserByIdRepositoryWithRightId(t *testing.T) {
	sut, repo := MakeDeleteUserSut()
	sut.Delete(FAKE_USER_ID)
	require.Equal(t, FAKE_USER_ID, repo.CheckByIdInput)
}

func TestDeleteUserUseCase_ShouldCallDelteUserRepositoryWithRightId(t *testing.T) {
	sut, repo := MakeDeleteUserSut()
	sut.Delete(FAKE_USER_ID)
	require.Equal(t, FAKE_USER_ID, repo.DeleteInput)
}

func TestDeleteUserUseCase_ShouldReturnErrorIfWrongIdIIsProvided(t *testing.T) {
	sut, repo := MakeDeleteUserSut()
	repo.CheckByIdOuput = false

	message, err := sut.Delete("any_wrong_id")

	require.Error(t, err)
	require.Equal(t, "No user with the provided id!", err.Error())
	require.Equal(t, "", message)
}

func TestDeleteUserUseCase_ShouldReturnErrorIfDelteRepositoryReturnsError(t *testing.T) {
	sut, repo := MakeDeleteUserSut()
	repo.DeleteOutput = errors.New("Internal error")

	message, err := sut.Delete(FAKE_USER_ID)

	require.Error(t, err)
	require.Equal(t, "", message)
	require.Equal(t, "Internal error", err.Error())
}
