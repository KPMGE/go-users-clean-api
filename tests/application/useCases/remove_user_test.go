package usecases_test

import (
	"errors"
	"testing"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

const FAKE_USER_ID string = "any_user_id"

func MakeDeleteUserSut() (*usecases.DeleteUserUseCase, *mocks_test.UserRepositorySpy) {
	repo := mocks_test.NewUserRepositorySpy()
	repo.DeleteOutput = nil
	repo.CheckByIdOuput = true
	sut := usecases.NewDeleteUserUseCase(repo)
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
