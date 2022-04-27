package controllers_test

import (
	"testing"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

type GetBookByIdController struct {
	useCase *usecases.GetBookByIdUseCase
}

func NewGetBookByIdController(useCase *usecases.GetBookByIdUseCase) *GetBookByIdController {
	return &GetBookByIdController{
		useCase: useCase,
	}
}

func (controller *GetBookByIdController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	controller.useCase.GetById(string(request.Params))
	return nil
}

func TestGetGiftByIdController_ShouldCallUseCaseWithRightData(t *testing.T) {
	repo := mocks_test.NewGetBookByIdRepositorySpy()
	useCase := usecases.NewGetBookByIdUseCase(repo)
	sut := NewGetBookByIdController(useCase)

	fakeRequest := protocols.NewHtppRequest(nil, []byte("any_book_id"))
	sut.Handle(fakeRequest)

	require.Equal(t, "any_book_id", repo.Input)
}
