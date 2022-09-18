package mocks_test

type TokenProviderStub struct {
	Output string
	Error  error
}

func (t *TokenProviderStub) Generate(data any) (string, error) {
	return t.Output, t.Error
}
