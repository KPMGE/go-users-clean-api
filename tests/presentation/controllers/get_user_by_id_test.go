package controllers_test

import (
	"errors"
	"testing"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

const FAKE_USER_ID string = "any_user_id"

type GetUserByIdController struct {
	useCase *usecases.GetUserByIdUseCase
}

func NewGetUserByIdController(useCase *usecases.GetUserByIdUseCase) *GetUserByIdController {
	return &GetUserByIdController{
		useCase: useCase,
	}
}

func (controller *GetUserByIdController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	if len(request.Params) == 0 {
		err := errors.New("Blank userId!")
		return helpers.BadRequest(err)
	}

	_, err := controller.useCase.Get(string(request.Params))
	if err != nil {
		return helpers.ServerError(err)
	}
	return nil
}

func MakeGetUserByIdController() (*GetUserByIdController, *mocks_test.UserRepositorySpy) {
	repo := mocks_test.NewUserRepositorySpy()
	useCase := usecases.NewGetUserByIdUseCase(repo)
	sut := NewGetUserByIdController(useCase)
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
