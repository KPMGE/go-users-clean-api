package controllermocks_test

import domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"

type LoginServiceMock struct {
	Input  *domaindto.LoginInputDTO
	Output string
	Error  error
}

func (s *LoginServiceMock) Login(input *domaindto.LoginInputDTO) (string, error) {
	s.Input = input
	return s.Output, s.Error
}

func NewLoginServiceMock() *LoginServiceMock {
	return &LoginServiceMock{
		Input:  nil,
		Output: "token",
		Error:  nil,
	}
}
