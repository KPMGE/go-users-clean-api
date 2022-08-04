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
	user.Books = []entities.Book{}
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
	var user entities.User

	result := repo.db.First(&user, fmt.Sprintf("user_name = '%s'", userName))

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}

	CheckError(result.Error)

	return true
}

func (repo *PostgresUserRepository) List() []*entities.User {
	var users []*entities.User
	var books []entities.Book

	resultUsers := repo.db.Find(&users)
	CheckError(resultUsers.Error)

	for _, user := range users {
		resultBooks := repo.db.Find(&books, fmt.Sprintf("user_id = '%s'", user.ID))
		CheckError(resultBooks.Error)
		user.Books = books
	}

	return users
}

func (repo *PostgresUserRepository) Delete(userId string) error {
	return nil
}

func (repo *PostgresUserRepository) CheckById(userId string) bool {
	var user entities.User

	result := repo.db.First(&user, fmt.Sprintf("id = '%s'", userId))

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}

	CheckError(result.Error)

	return true
}

func (repo *PostgresUserRepository) GetById(userId string) (*entities.User, error) {
	var user entities.User

	result := repo.db.First(&user, fmt.Sprintf("id = '%s'", userId))

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	CheckError(result.Error)

	return &user, nil
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}
