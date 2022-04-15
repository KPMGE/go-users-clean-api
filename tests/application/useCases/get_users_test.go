package usecases_test

import (
	"testing"

	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

type ListUsersUseCase struct {
	userRepository protocols.UserRepository
}

func NewListUsersUseCase(repo protocols.UserRepository) *ListUsersUseCase {
	return &ListUsersUseCase{
		userRepository: repo,
	}
}

func (useCase *ListUsersUseCase) List() []*dto.ListUsersDTO {
	users := useCase.userRepository.List()
	return dto.MapListUsersDTO(users)
}

func makeListUsersSut() (*ListUsersUseCase, *mocks_test.UserRepositorySpy) {
	repo := mocks_test.NewUserRepositorySpy()
	repo.ListUsersOutput = []*entities.User{}
	sut := NewListUsersUseCase(repo)
	return sut, repo
}

func TestListUsersUseCase_WhenRepositoryReturnsBlankArray(t *testing.T) {
	sut, _ := makeListUsersSut()
	users := sut.List()

	require.Equal(t, []*dto.ListUsersDTO{}, users)
}

func TestListUsersUseCase_WhenRepositoryReturnsFilledArray(t *testing.T) {
	sut, repo := makeListUsersSut()
	fakeUser, _ := entities.NewUser("any_name", "any_username", "any_valid_email@gmail.com")
	repo.ListUsersOutput = []*entities.User{fakeUser}
	fakeList := dto.MapListUsersDTO([]*entities.User{fakeUser})

	users := sut.List()

	require.ElementsMatch(t, fakeList, users)
}
