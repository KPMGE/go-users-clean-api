package repositories

import (
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
)

var accounts []*entities.Account

type InMemoryAccountRepository struct{}

func (repo *InMemoryAccountRepository) CheckAccountByEmail(email string) bool {
	for _, account := range accounts {
		if account.Email == email {
			return true
		}
	}
	return false
}

func (repo *InMemoryAccountRepository) CheckAccountByUserName(userName string) bool {
	for _, account := range accounts {
		if account.UserName == userName {
			return true
		}
	}
	return false
}

func (repo *InMemoryAccountRepository) Save(account *entities.Account) error {
	accounts = append(accounts, account)
	return nil
}

func (repo *InMemoryAccountRepository) DeleteAccountById(accountId string) bool {
	for index, account := range accounts {
		if account.ID == accountId {
			accounts = removeIndex(accounts, index)
			return true
		}
	}
	return false
}

func NewInmemoryAccountRepository() *InMemoryAccountRepository {
	return &InMemoryAccountRepository{}
}
