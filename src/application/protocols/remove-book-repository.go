package protocols

type RemoveBookRepository interface {
	Remove(bookId string) error
}
