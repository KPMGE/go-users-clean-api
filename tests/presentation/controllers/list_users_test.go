package controllers_test

import (
	"encoding/json"
	"fmt"
	"testing"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	"github.com/stretchr/testify/require"
)

type ListUsersController struct {
	useCase *usecases.ListUsersUseCase
}

func NewListUsersController(useCase *usecases.ListUsersUseCase) *ListUsersController {
	return &ListUsersController{
		useCase: useCase,
	}
}

func (controller *ListUsersController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	users := controller.useCase.List()
	jsonUsers, err := json.Marshal(users)
	fmt.Println(users)
	if err != nil {
		panic(err)
	}
	return helpers.Ok(jsonUsers)
}

func TestListUsersController_WithNoUsers(t *testing.T) {
	repo := repositories.NewInMemoryUserRepository()
	useCase := usecases.NewListUsersUseCase(repo)
	sut := NewListUsersController(useCase)

	fakeRequest := protocols.NewHtppRequest(nil, nil)
	httpResponse := sut.Handle(fakeRequest)

	require.Equal(t, 200, httpResponse.StatusCode)
}
