package usecases_test

import (
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
	useCase.userRepository.CheckById(userId)
	useCase.userRepository.Delete(userId)
	return "user deleted successfully", nil
}

func MakeDeleteUserSut() (*DeleteUserUseCase, *mocks_test.UserRepositorySpy) {
	repo := mocks_test.NewUserRepositorySpy()
	repo.DeleteOutput = nil
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
