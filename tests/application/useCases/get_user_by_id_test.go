package usecases_test

import (
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

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

func (useCase *GetUserByIdUseCase) Get(userId string) {
	useCase.userRepository.GetById(userId)
}

func TestGetUserByIdUseCase_ShouldCallRepositoryCorrectly(t *testing.T) {
	repo := mocks_test.NewUserRepositorySpy()

	sut := NewGetUserByIdUseCase(repo)

	sut.Get(FAKE_USER_ID)
	require.Equal(t, FAKE_USER_ID, repo.GetByidInput)
}
