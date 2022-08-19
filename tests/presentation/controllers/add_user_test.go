package controllers_test

import (
	"encoding/json"
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/application/services"
	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	"github.com/stretchr/testify/require"
)

const fakeName string = "any_name"
const fakeUserName string = "any_username"
const fakeEmail string = "any_valid_email@gmail.com"

func makeFakeAddUserRequest(name string, userName string, email string) *protocols.HttpRequest {
	input := domaindto.NewAddUserInputDTO(name, userName, email)
	jsonEntry, err := json.Marshal(input)

	if err != nil {
		panic("Error generating json")
	}

	return protocols.NewHtppRequest(jsonEntry, nil)
}

func makeAddUserControllerSut() *controllers.AddUserController {
	repo := repositories.NewInMemoryUserRepository()
	service := services.NewAddUserService(repo)
	sut := controllers.NewAddUserController(service)
	return sut
}

func convertJsonToAccoutOutputDTO(data []byte) *domaindto.AddAccountOutputDTO {
	var bodyObj domaindto.AddAccountOutputDTO
	err := json.Unmarshal(data, &bodyObj)
	if err != nil {
		panic(err)
	}
	return &bodyObj
}

func TestAdduserController_WithCorrectInput(t *testing.T) {
	sut := makeAddUserControllerSut()
	fakeRequest := makeFakeAddUserRequest(fakeName, fakeUserName, fakeEmail)

	httpResponse := sut.Handle(fakeRequest)
	bodyObj := convertJsonToAccoutOutputDTO(httpResponse.JsonBody)

	require.Equal(t, 200, httpResponse.StatusCode)
	require.NotNil(t, bodyObj.ID)
	require.Equal(t, bodyObj.Email, fakeEmail)
	require.Equal(t, bodyObj.UserName, fakeUserName)
}

func TestAdduserController_WithInvalidEmail(t *testing.T) {
	sut := makeAddUserControllerSut()
	fakeRequest := makeFakeAddUserRequest(fakeName, fakeUserName, "invalid email")

	httpResponse := sut.Handle(fakeRequest)

	require.Equal(t, 400, httpResponse.StatusCode)
	require.Equal(t, "Invalid email!", string(httpResponse.JsonBody))
}

func TestAdduserController_WithBlankFields(t *testing.T) {
	sut := makeAddUserControllerSut()

	fakeRequest := makeFakeAddUserRequest("", fakeUserName, fakeEmail)
	httpResponse := sut.Handle(fakeRequest)
	require.Equal(t, 400, httpResponse.StatusCode)
	require.NotNil(t, httpResponse.JsonBody)

	fakeRequest = makeFakeAddUserRequest(fakeName, "", fakeEmail)
	httpResponse = sut.Handle(fakeRequest)
	require.Equal(t, 400, httpResponse.StatusCode)
	require.NotNil(t, httpResponse.JsonBody)

	fakeRequest = makeFakeAddUserRequest(fakeName, fakeUserName, "")
	httpResponse = sut.Handle(fakeRequest)
	require.Equal(t, 400, httpResponse.StatusCode)
	require.NotNil(t, httpResponse.JsonBody)
}

func TestAdduserController_WithInvalidJsonInput(t *testing.T) {
	sut := makeAddUserControllerSut()
	fakeRequest := protocols.NewHtppRequest([]byte("invalid json"), nil)

	httpResponse := sut.Handle(fakeRequest)

	require.Equal(t, 500, httpResponse.StatusCode)
	require.NotNil(t, httpResponse.JsonBody)
}
