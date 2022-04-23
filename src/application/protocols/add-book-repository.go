package protocols

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type AddBookRepository interface {
	Add(newBook *entities.Book) (*entities.Book, error)
}
