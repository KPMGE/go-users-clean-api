package protocols

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type ListBooksRepository interface {
	List() ([]*entities.Book, error)
}
