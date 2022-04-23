package mocks_test

type RemoveBookRepositorySpy struct {
	RemoveInput string
	RemoveError error
}

func (repo *RemoveBookRepositorySpy) Remove(bookId string) error {
	repo.RemoveInput = bookId
	return repo.RemoveError
}

func NewRemoveBookRepositorySpy() *RemoveBookRepositorySpy {
	return &RemoveBookRepositorySpy{}
}
