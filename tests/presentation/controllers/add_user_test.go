package controllers_test

import (
	"errors"
	"net/http"
	"testing"

	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	controllermocks_test "github.com/KPMGE/go-users-clean-api/tests/presentation/controller-mocks"
	"github.com/stretchr/testify/require"
)

const fakeName string = "any_name"
const fakeUserName string = "any_username"
const fakeEmail string = "any_valid_email@gmail.com"

type AddUserServiceMock struct {
	Output *domaindto.AddUserOutputDTO
	Error  error
}

func (s *AddUserServiceMock) Add(input *domaindto.AddUserInputDTO) (*domaindto.AddUserOutputDTO, error) {
	return s.Output, s.Error
}

func MakeFakeAddUserServiceMock() *AddUserServiceMock {
	return &AddUserServiceMock{
		Output: domaindto.NewAddUserOutputDTO("any_id", "any_name", "any_username", "any_valid_email@gmail.com"),
		Error:  nil,
	}
}

func makeAddUserControllerSut() (*controllers.AddUserController, *AddUserServiceMock, *controllermocks_test.ValidatorMock) {
	serviceMock := MakeFakeAddUserServiceMock()
	validatorMock := controllermocks_test.ValidatorMock{Output: nil}
	sut := controllers.NewAddUserController(serviceMock, &validatorMock)
	return sut, serviceMock, &validatorMock
}

func makeFakeAddUserRequest() *protocols.HttpRequest {
	fakeBody := `{
		"name":     "any_name",
		"userName": "any_username",
		"email":    "any_email@gmail.com"
	}`

	fakeRequest := protocols.NewHttpRequest([]byte(fakeBody), nil)
	return fakeRequest
}

func TestAdduserController_ShouldReturnRightDataOnSuccess(t *testing.T) {
	sut, serviceMock, _ := makeAddUserControllerSut()

	httpResponse := sut.Handle(makeFakeAddUserRequest())

	require.Equal(t, http.StatusOK, httpResponse.StatusCode)
	require.Equal(t, serviceMock.Output, httpResponse.Body)
}

func TestAdduserController_ShouldReturnBadRequestIfValidatorReturnsError(t *testing.T) {
	sut, _, validatorMock := makeAddUserControllerSut()
	validatorMock.Output = errors.New("validation error")

	httpResponse := sut.Handle(makeFakeAddUserRequest())

	require.Equal(t, http.StatusBadRequest, httpResponse.StatusCode)
	require.Equal(t, validatorMock.Output.Error(), httpResponse.Body)
}
