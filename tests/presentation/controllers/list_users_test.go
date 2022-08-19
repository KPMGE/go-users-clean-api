package controllers_test

import (
	"encoding/json"
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/application/services"
	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

func MakeListUsersSut() (*controllers.ListUsersController, *mocks_test.UserRepositorySpy) {
	repo := mocks_test.NewUserRepositorySpy()
	repo.ListUsersOutput = []*entities.User{}
	service := services.NewListUsersService(repo)
	sut := controllers.NewListUsersController(service)
	return sut, repo
}

func TestListUsersController_WithNoUsers(t *testing.T) {
	sut, _ := MakeListUsersSut()
	fakeRequest := protocols.NewHtppRequest(nil, nil)
	httpResponse := sut.Handle(fakeRequest)

	var listObj []*domaindto.ListUsersDTO
	json.Unmarshal(httpResponse.JsonBody, &listObj)

	require.Equal(t, 200, httpResponse.StatusCode)
	require.Equal(t, 0, len(listObj))
	require.Equal(t, []*domaindto.ListUsersDTO{}, listObj)
}

func TestListUsersController_WithTwoUsers(t *testing.T) {
	sut, repo := MakeListUsersSut()

	fakeUser, _ := entities.NewUser("any_name", "any_username", "any_valid_email@gmail.com")
	repo.ListUsersOutput = []*entities.User{fakeUser, fakeUser}

	fakeRequest := protocols.NewHtppRequest(nil, nil)
	httpResponse := sut.Handle(fakeRequest)

	var objBody []*domaindto.ListUsersDTO
	err := json.Unmarshal(httpResponse.JsonBody, &objBody)
	if err != nil {
		panic(err)
	}

	require.Equal(t, 200, httpResponse.StatusCode)
	require.Equal(t, 2, len(objBody))
}
