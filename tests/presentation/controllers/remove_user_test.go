package controllers_test

import (
	"testing"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

const FAKE_OUT_MESSAGE string = "any_message_from_usecase"

func MakeDeleteUserControllerSut() (*controllers.DeleteUserController, *mocks_test.UserRepositorySpy) {
	repo := mocks_test.NewUserRepositorySpy()
	repo.CheckByIdOuput = true
	useCase := usecases.NewDeleteUserUseCase(repo)
	sut := controllers.NewDeleteUserController(useCase)
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
