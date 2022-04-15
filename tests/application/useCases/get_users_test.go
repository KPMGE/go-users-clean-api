package usecases_test

import (
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

type ListUsersDTO struct {
	ID       string
	Name     string
	UserName string
	Email    string
}

func NewListUserDTO(id string, name string, userName string, email string) *ListUsersDTO {
	return &ListUsersDTO{
		ID:       id,
		Name:     name,
		UserName: userName,
		Email:    email,
	}
}

func MapListUsersDTO(users []*entities.User) []*ListUsersDTO {
	list := []*ListUsersDTO{}
	for _, user := range users {
		list = append(list, NewListUserDTO(user.ID, user.Name, user.UserName, user.Email))
	}
	return list
}

type ListUsersUseCase struct {
	userRepository protocols.UserRepository
}

func NewListUsersUseCase(repo protocols.UserRepository) *ListUsersUseCase {
	return &ListUsersUseCase{
		userRepository: repo,
	}
}

func (useCase *ListUsersUseCase) List() []*ListUsersDTO {
	users := useCase.userRepository.List()
	return MapListUsersDTO(users)
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

	require.Equal(t, []*ListUsersDTO{}, users)
}

func TestListUsersUseCase_WhenRepositoryReturnsFilledArray(t *testing.T) {
	sut, repo := makeListUsersSut()
	fakeUser, _ := entities.NewUser("any_name", "any_username", "any_valid_email@gmail.com")
	repo.ListUsersOutput = []*entities.User{fakeUser}
	fakeList := MapListUsersDTO([]*entities.User{fakeUser})

	users := sut.List()

	require.ElementsMatch(t, fakeList, users)
}
