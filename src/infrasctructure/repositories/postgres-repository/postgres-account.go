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
		log.Fatalln(err)
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
	result := repo.db.Delete(&entities.Account{}, fmt.Sprintf("id = '%s'", accountId))
	CheckError(result.Error)
	if result.RowsAffected < 1 {
		return false
	}
	return true
}

func (repo *PostgresAccountRepository) ListAccounts() []entities.Account {
	var accounts []entities.Account

	result := repo.db.Find(&accounts)
	CheckError(result.Error)

	return accounts
}

func NewPostgresAccountRepository(db *gorm.DB) *PostgresAccountRepository {
	return &PostgresAccountRepository{
		db: db,
	}
}
