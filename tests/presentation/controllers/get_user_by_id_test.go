package controllers_test

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
	"testing"
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
	controller.useCase.Get(string(request.Params))
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
