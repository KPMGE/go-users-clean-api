package controllers_test

import (
	"net/http"
	"testing"

	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"github.com/stretchr/testify/require"
)

type ListUsersServiceMock struct {
	Output []*domaindto.ListUsersDTO
}

func (l *ListUsersServiceMock) List() []*domaindto.ListUsersDTO {
	return l.Output
}

func MakeListUsersSut() (*controllers.ListUsersController, *ListUsersServiceMock) {
	serviceMock := ListUsersServiceMock{Output: []*domaindto.ListUsersDTO{}}
	sut := controllers.NewListUsersController(&serviceMock)
	return sut, &serviceMock
}

func TestListUsersController_ShouldReturnFromService(t *testing.T) {
	sut, serviceMock := MakeListUsersSut()

	httpResponse := sut.Handle(nil)

	require.Equal(t, http.StatusOK, httpResponse.StatusCode)
	require.NotNil(t, serviceMock.Output, httpResponse.Body)
}
