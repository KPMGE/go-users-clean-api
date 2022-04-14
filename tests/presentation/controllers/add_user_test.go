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
	return helpers.Ok(output)
}

func NewAddUserController(useCase *usecases.AddUserUseCase) *AddUserController {
	return &AddUserController{
		useCase: useCase,
	}
}

func makeFakeValidRequest() *HttpRequest {
	input := dto.NewAddUserInputDTO("any_name", "any_username", "any_valid_email@gmail.com")
	jsonEntry, err := json.Marshal(input)

	if err != nil {
		panic("Error generating json")
	}

	return NewHtppRequest(jsonEntry, nil)
}

func TestAdduserController_WithCorrectInput(t *testing.T) {
	repo := repositories.NewInMemoryUserRepository()
	useCase := usecases.NewAddUserUseCase(repo)
	sut := NewAddUserController(useCase)
	fakeRequest := makeFakeValidRequest()

	httpResponse := sut.Handle(fakeRequest)

	require.Equal(t, 200, httpResponse.StatusCode)
	require.NotNil(t, httpResponse.Body)
}
