package controllers_test

import (
	"testing"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

type RemoveBookController struct {
	useCase *usecases.RemoveBookUseCase
}

func NewRemoveBookController(useCase *usecases.RemoveBookUseCase) *RemoveBookController {
	return &RemoveBookController{
		useCase: useCase,
	}
}

func (controller *RemoveBookController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	controller.useCase.Remove(string(request.Params))
	return nil
}

func TestRemoveBookController_ShoulCallUseCaseWithRightBookId(t *testing.T) {
	removeBookRepo := mocks_test.NewRemoveBookRepositorySpy()
	findBookRepo := mocks_test.NewFindBookRepositorySpy()
	fakeBook, _ := entities.NewBook("any_title", "any_author", "any_description", 100, "any_user_id")
	findBookRepo.FindOutput = fakeBook
	useCase := usecases.NewRemoveBookUseCase(removeBookRepo, findBookRepo)
	sut := NewRemoveBookController(useCase)

	fakeRequest := protocols.NewHtppRequest(nil, []byte("any_book_id"))
	sut.Handle(fakeRequest)

	require.Equal(t, "any_book_id", findBookRepo.FindInput)
	require.Equal(t, "any_book_id", removeBookRepo.RemoveInput)
}
