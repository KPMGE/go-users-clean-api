package controllers_test

import (
	"testing"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

const FAKE_OUT_MESSAGE string = "any_message_from_usecase"

type DeleteUserController struct {
	useCase *usecases.DeleteUserUseCase
}

func NewDeleteUserController(useCase *usecases.DeleteUserUseCase) *DeleteUserController {
	return &DeleteUserController{
		useCase: useCase,
	}
}

func (controller *DeleteUserController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	return helpers.Ok([]byte("user deleted successfully"))
}

func MakeDeleteUserControllerSut() *DeleteUserController {
	repo := mocks_test.NewUserRepositorySpy()
	useCase := usecases.NewDeleteUserUseCase(repo)
	sut := NewDeleteUserController(useCase)
	return sut
}

func TestDeleteUserController_ShouldCallUseCaseWithRightData(t *testing.T) {
	sut := MakeDeleteUserControllerSut()
	fakeRequest := protocols.NewHtppRequest(nil, []byte("any_valid_id"))
	httpResponse := sut.Handle(fakeRequest)
	require.Equal(t, 200, httpResponse.StatusCode)
	require.Equal(t, "user deleted successfully", string(httpResponse.JsonBody))
}
