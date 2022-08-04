package postgresrepository

import (
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"gorm.io/gorm"
)

type PostgresBookRepository struct {
	db *gorm.DB
}

func (repo *PostgresBookRepository) List() ([]*entities.Book, error) {
	return nil, nil
}

func (repo *PostgresBookRepository) Add(newBook *entities.Book) (*entities.Book, error) {
	result := repo.db.Create(newBook)
	if result.Error != nil {
		return nil, result.Error
	}
	return newBook, nil
}

func (repo *PostgresBookRepository) Find(bookId string) (*entities.Book, error) {
	return nil, nil
}

func (repo *PostgresBookRepository) Remove(bookId string) error {
	return nil
}

func (repo *PostgresBookRepository) Get(bookId string) (*entities.Book, error) {
	return nil, nil
}

func NewPostgresBookRepository(db *gorm.DB) *PostgresBookRepository {
	return &PostgresBookRepository{
		db: db,
	}
}
