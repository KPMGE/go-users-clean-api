package postgresrepository

import (
	"errors"
	"fmt"
	"log"

	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"gorm.io/gorm"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type PostgresAccountRepository struct {
	db *gorm.DB
}

func (repo *PostgresAccountRepository) CheckAccountByEmail(email string) bool {
	var account entities.Account

	result := repo.db.First(&account, fmt.Sprintf("email='%s'", email))

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}

	CheckError(result.Error)

	return true
}

func (repo *PostgresAccountRepository) CheckAccountByUserName(userName string) bool {
	var account entities.Account

	result := repo.db.First(&account, fmt.Sprintf("user_name = '%s'", userName))

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}

	CheckError(result.Error)

	return true
}

func (repo *PostgresAccountRepository) Save(account *entities.Account) error {
	result := repo.db.Create(account)
	CheckError(result.Error)
	return nil
}

func (repo *PostgresAccountRepository) DeleteAccountById(accountId string) bool {
	return true
}

func NewPostgresAccountRepository(db *gorm.DB) *PostgresAccountRepository {
	return &PostgresAccountRepository{
		db: db,
	}
}
