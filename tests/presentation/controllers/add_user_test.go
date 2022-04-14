package controllers_test

import (
	"encoding/json"
	"testing"

	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	"github.com/stretchr/testify/require"
)

type HttpRequest struct {
	Params []byte
	Body   []byte
}

func NewHtppRequest(body []byte, params []byte) *HttpRequest {
	return &HttpRequest{
		Body:   body,
		Params: params,
	}
}

type AddUserController struct {
	useCase *usecases.AddUserUseCase
}

func (controller *AddUserController) Handle(request *HttpRequest) *protocols.HttpResponse {
	var inputUser dto.AddUserInputDTO
	err := json.Unmarshal(request.Body, &inputUser)
	if err != nil {
		panic(err)
	}

	newUser, err := entities.NewUser(inputUser.Name, inputUser.UserName, inputUser.Email)
	if err != nil {
		return nil
	}

	output := dto.NewAddUserOutputDTO(newUser.ID, newUser.Name, newUser.UserName, newUser.Email)
	outputJson, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}
	return helpers.Ok(string(outputJson))
}

func NewAddUserController(useCase *usecases.AddUserUseCase) *AddUserController {
	return &AddUserController{
		useCase: useCase,
	}
}

const fakeName string = "any_name"
const fakeUserName string = "any_username"
const fakeEmail string = "any_valid_email@gmail.com"

func makeFakeValidRequest() *HttpRequest {
	input := dto.NewAddUserInputDTO(fakeName, fakeUserName, fakeEmail)
	jsonEntry, err := json.Marshal(input)

	if err != nil {
		panic("Error generating json")
	}

	return NewHtppRequest(jsonEntry, nil)
}

func makeAddUserControllerSut() *AddUserController {
	repo := repositories.NewInMemoryUserRepository()
	useCase := usecases.NewAddUserUseCase(repo)
	sut := NewAddUserController(useCase)
	return sut
}

func convertJsonToAccoutOutputDTO(data string) *dto.AddAccountOutputDTO {
	var bodyObj dto.AddAccountOutputDTO
	err := json.Unmarshal([]byte(data), &bodyObj)
	if err != nil {
		panic(err)
	}
	return &bodyObj
}

func TestAdduserController_WithCorrectInput(t *testing.T) {
	sut := makeAddUserControllerSut()
	fakeRequest := makeFakeValidRequest()
	httpResponse := sut.Handle(fakeRequest)
	bodyObj := convertJsonToAccoutOutputDTO(httpResponse.JsonBody)

	require.Equal(t, 200, httpResponse.StatusCode)
	require.NotNil(t, bodyObj.ID)
	require.Equal(t, bodyObj.Email, fakeEmail)
	require.Equal(t, bodyObj.UserName, fakeUserName)
}
