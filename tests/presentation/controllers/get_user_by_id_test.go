package controllers_test

import (
	"encoding/json"
	"errors"
	"log"
	"testing"

	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

const FAKE_USER_ID string = "any_user_id"

func MakeGetUserByIdController() (*controllers.GetUserByIdController, *mocks_test.UserRepositorySpy) {
	repo := mocks_test.NewUserRepositorySpy()
	useCase := usecases.NewGetUserByIdUseCase(repo)
	sut := controllers.NewGetUserByIdController(useCase)
	return sut, repo
}

func TestGetUserByIdController_ShouldCallUseCaseWithRightData(t *testing.T) {
	sut, repo := MakeGetUserByIdController()
	fakeRequest := protocols.NewHtppRequest(nil, []byte(FAKE_USER_ID))

	sut.Handle(fakeRequest)

	require.Equal(t, FAKE_USER_ID, repo.GetByidInput)
}

func TestGetUserByIdController_ShouldReturnErrorIfParamsIsBlank(t *testing.T) {
	sut, _ := MakeGetUserByIdController()
	fakeRequest := protocols.NewHtppRequest(nil, []byte(""))

	httpResponse := sut.Handle(fakeRequest)

	require.Equal(t, 400, httpResponse.StatusCode)
	require.Equal(t, "Blank userId!", string(httpResponse.JsonBody))
}

func TestGetUserByIdController_ShouldReturnErrorIfUseCaseReturnsError(t *testing.T) {
	sut, repo := MakeGetUserByIdController()
	repo.GetByidError = errors.New("some server error")
	fakeRequest := protocols.NewHtppRequest(nil, []byte(FAKE_USER_ID))

	httpResponse := sut.Handle(fakeRequest)

	require.Equal(t, 500, httpResponse.StatusCode)
	require.Equal(t, "some server error", string(httpResponse.JsonBody))
}

func TestGetUserByIdController_ShouldReturnDataOnSuccess(t *testing.T) {
	sut, repo := MakeGetUserByIdController()
	fakeUser, _ := entities.NewUser(fakeName, fakeUserName, fakeEmail)
	repo.GetByidOutput = fakeUser
	fakeRequest := protocols.NewHtppRequest(nil, []byte(FAKE_USER_ID))

	httpResponse := sut.Handle(fakeRequest)

	// convert json to struct
	var outputObj *dto.GetUserByIdUseCaseOutputDTO
	err := json.Unmarshal(httpResponse.JsonBody, &outputObj)
	if err != nil {
		log.Fatal(err)
	}

	require.Equal(t, 200, httpResponse.StatusCode)
	require.Equal(t, fakeUser.ID, outputObj.ID)
	require.Equal(t, fakeUser.Email, outputObj.Email)
	require.Equal(t, fakeUser.Name, outputObj.Name)
	require.Equal(t, fakeUser.UserName, outputObj.UserName)
}
