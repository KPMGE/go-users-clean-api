package controllers_test

import (
	"encoding/json"
	"testing"

	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

func MakeListUsersSut() (*controllers.ListUsersController, *mocks_test.UserRepositorySpy) {
	repo := mocks_test.NewUserRepositorySpy()
	repo.ListUsersOutput = []*entities.User{}
	useCase := usecases.NewListUsersUseCase(repo)
	sut := controllers.NewListUsersController(useCase)
	return sut, repo
}

func TestListUsersController_WithNoUsers(t *testing.T) {
	sut, _ := MakeListUsersSut()
	fakeRequest := protocols.NewHtppRequest(nil, nil)
	httpResponse := sut.Handle(fakeRequest)

	var listObj []*dto.ListUsersDTO
	json.Unmarshal(httpResponse.JsonBody, &listObj)

	require.Equal(t, 200, httpResponse.StatusCode)
	require.Equal(t, 0, len(listObj))
	require.Equal(t, []*dto.ListUsersDTO{}, listObj)
}

func TestListUsersController_WithTwoUsers(t *testing.T) {
	sut, repo := MakeListUsersSut()

	fakeUser, _ := entities.NewUser("any_name", "any_username", "any_valid_email@gmail.com")
	repo.ListUsersOutput = []*entities.User{fakeUser, fakeUser}

	fakeRequest := protocols.NewHtppRequest(nil, nil)
	httpResponse := sut.Handle(fakeRequest)

	var objBody []*dto.ListUsersDTO
	err := json.Unmarshal(httpResponse.JsonBody, &objBody)
	if err != nil {
		panic(err)
	}

	require.Equal(t, 200, httpResponse.StatusCode)
	require.Equal(t, 2, len(objBody))
}
