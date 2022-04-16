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
	message, err := controller.useCase.Delete(string(request.Params))
	if err != nil {
		return helpers.BadRequest(err)
	}
	return helpers.Ok([]byte(message))
}

func MakeDeleteUserControllerSut() (*DeleteUserController, *mocks_test.UserRepositorySpy) {
	repo := mocks_test.NewUserRepositorySpy()
	repo.CheckByIdOuput = true
	useCase := usecases.NewDeleteUserUseCase(repo)
	sut := NewDeleteUserController(useCase)
	return sut, repo
}

func TestDeleteUserController_WhenCalledWithRightData(t *testing.T) {
	sut, _ := MakeDeleteUserControllerSut()
	fakeRequest := protocols.NewHtppRequest(nil, []byte("any_valid_id"))

	httpResponse := sut.Handle(fakeRequest)

	require.Equal(t, 200, httpResponse.StatusCode)
	require.Equal(t, "user deleted successfully", string(httpResponse.JsonBody))
}

func TestDeleteUserController_WhenCalledWithWrongData(t *testing.T) {
	sut, repo := MakeDeleteUserControllerSut()
	repo.CheckByIdOuput = false

	fakeRequest := protocols.NewHtppRequest(nil, []byte("any_invalid_id"))

	httpResponse := sut.Handle(fakeRequest)

	require.Equal(t, 400, httpResponse.StatusCode)
	require.Equal(t, "No user with the provided id!", string(httpResponse.JsonBody))
}
