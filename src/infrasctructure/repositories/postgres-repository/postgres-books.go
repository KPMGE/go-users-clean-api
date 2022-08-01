package postgresrepository

import (
	"database/sql"

	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
)

type PostgresBookRepository struct {
	db *sql.DB
}

func removeIndex[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

func (repo *PostgresBookRepository) List() ([]*entities.Book, error) {
	return nil, nil
}

func (repo *PostgresBookRepository) Add(newBook *entities.Book) (*entities.Book, error) {
	query := `INSERT INTO "books"(
  "id", "created_at", "updated_at", "title", "author", "price", "description", "user_id") 
  VALUES($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := repo.db.Exec(query, newBook.ID, newBook.CreatedAt, newBook.UpdatedAt,
		newBook.Title, newBook.Author, newBook.Price, newBook.Description, newBook.UserId)

	if err != nil {
		return nil, err
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

func NewPostgresBookRepository(db *sql.DB) *PostgresBookRepository {
	return &PostgresBookRepository{
		db: db,
	}
}
