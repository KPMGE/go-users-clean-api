package postgresrepository

import (
	"database/sql"

	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func (repo *PostgresUserRepository) Save(user *entities.User) error {
	query := `INSERT INTO "users"("id", "created_at", "updated_at", "name", "user_name", "email") VALUES($1, $2, $3, $4, $5, $6)`
	_, err := repo.db.Exec(query, user.ID, user.CreatedAt, user.UpdatedAt, user.Name, user.UserName, user.Email)

	return err
}

func (repo *PostgresUserRepository) CheckByEmail(email string) bool {
	return false
}

func (repo *PostgresUserRepository) CheckByUserName(userName string) bool {
	return false
}

func (repo *PostgresUserRepository) List() []*entities.User {
	return nil
}

func (repo *PostgresUserRepository) Delete(userId string) error {
	return nil
}

func (repo *PostgresUserRepository) CheckById(userId string) bool {
	return false
}

func (repo *PostgresUserRepository) GetById(userId string) (*entities.User, error) {
	return nil, nil
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}
