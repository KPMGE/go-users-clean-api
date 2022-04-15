package usecases_test

import (
	"testing"

	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

func makeListUsersSut() (*usecases.ListUsersUseCase, *mocks_test.UserRepositorySpy) {
	repo := mocks_test.NewUserRepositorySpy()
	repo.ListUsersOutput = []*entities.User{}
	sut := usecases.NewListUsersUseCase(repo)
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
