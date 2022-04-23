package protocols

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type FindBookRepository interface {
	Find(bookId string) (*entities.Book, error)
}
