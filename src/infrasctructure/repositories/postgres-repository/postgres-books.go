package postgresrepository

import (
	"errors"
	"fmt"

	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"gorm.io/gorm"
)

type PostgresBookRepository struct {
	db *gorm.DB
}

func (repo *PostgresBookRepository) List() ([]*entities.Book, error) {
	var books []*entities.Book

	result := repo.db.Find(&books)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return []*entities.Book{}, nil
	}

	CheckError(result.Error)

	return books, nil
}

func (repo *PostgresBookRepository) Add(newBook *entities.Book) (*entities.Book, error) {
	result := repo.db.Create(newBook)
	if result.Error != nil {
		return nil, result.Error
	}
	return newBook, nil
}

func (repo *PostgresBookRepository) Find(bookId string) (*entities.Book, error) {
	var foundBook entities.Book

	result := repo.db.First(&foundBook, fmt.Sprintf("id = '%s'", bookId))

	if result.Error != nil {
		return nil, result.Error
	}

	return &foundBook, nil
}

func (repo *PostgresBookRepository) Get(bookId string) (*entities.Book, error) {
	var foundBook entities.Book

	result := repo.db.First(&foundBook, fmt.Sprintf("id = '%s'", bookId))

	if result.Error != nil {
		return nil, result.Error
	}

	return &foundBook, nil
}

func (repo *PostgresBookRepository) Remove(bookId string) error {
	result := repo.db.Delete(&entities.Book{}, fmt.Sprintf("id = '%s'", bookId))
	return result.Error
}

func NewPostgresBookRepository(db *gorm.DB) *PostgresBookRepository {
	return &PostgresBookRepository{
		db: db,
	}
}
