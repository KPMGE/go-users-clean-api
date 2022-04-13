package mocks_test

type HasherSpy struct {
	Input string
}

func (hasher *HasherSpy) Hash(plainText string) string {
	hasher.Input = plainText
	return "hashed_text"
}

func NewHasherSpy() *HasherSpy {
	return &HasherSpy{}
}
