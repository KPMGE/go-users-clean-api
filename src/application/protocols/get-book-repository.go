package protocols

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type GetBookRepository interface {
	Get(bookId string) (*entities.Book, error)
}
