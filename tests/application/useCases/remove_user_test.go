package usecases_test

import (
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

type DeleteUserUseCase struct {
	userRepository protocols.UserRepository
}

func NewDeleteUserUseCase(repo protocols.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		userRepository: repo,
	}
}

func (useCase *DeleteUserUseCase) Delete(userId string) (string, error) {
	return "user deleted successfully", nil
}

func MakeDeleteUserSut() (*DeleteUserUseCase, *mocks_test.UserRepositorySpy) {
	repo := mocks_test.NewUserRepositorySpy()
	sut := NewDeleteUserUseCase(repo)
	return sut, repo
}

func TestDeleteUserUseCase_ShouldReturnRightDataOnSuccess(t *testing.T) {
	sut, _ := MakeDeleteUserSut()
	message, err := sut.Delete("any_right_id")

	require.Nil(t, err)
	require.Equal(t, "user deleted successfully", message)
}
