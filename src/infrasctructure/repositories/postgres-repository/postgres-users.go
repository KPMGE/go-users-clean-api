package postgresrepository

import (
	"errors"
	"fmt"

	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

func (repo *PostgresUserRepository) Save(user *entities.User) error {
	result := repo.db.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *PostgresUserRepository) CheckByEmail(email string) bool {
	var user entities.User

	result := repo.db.First(&user, fmt.Sprintf("email='%s'", email))

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}

	CheckError(result.Error)

	return true
}

func (repo *PostgresUserRepository) CheckByUserName(userName string) bool {
	return false
}

func (repo *PostgresUserRepository) List() []*entities.User {
	var users []*entities.User
	result := repo.db.Find(&users)
	CheckError(result.Error)
	return users
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

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}
