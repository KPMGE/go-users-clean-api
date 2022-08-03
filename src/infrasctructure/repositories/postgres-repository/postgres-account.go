package postgresrepository

import (
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
	return false
}

func (repo *PostgresAccountRepository) CheckAccountByUserName(userName string) bool {
	return false
}

func (repo *PostgresAccountRepository) Save(account *entities.Account) error {
	result := repo.db.Create(account)
	if result.Error != nil {
		return result.Error
	}
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
