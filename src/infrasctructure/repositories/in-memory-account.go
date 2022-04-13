package repositories

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type InMemoryAccountRepository struct {
	accounts []*entities.Account
}

func (repo *InMemoryAccountRepository) CheckAccountByEmail(email string) bool {
	for _, account := range repo.accounts {
		if account.Email == email {
			return true
		}
	}
	return false
}

func (repo *InMemoryAccountRepository) CheckAccountByUserName(userName string) bool {
	for _, account := range repo.accounts {
		if account.UserName == userName {
			return true
		}
	}
	return false
}

func (repo *InMemoryAccountRepository) Save(account *entities.Account) error {
	repo.accounts = append(repo.accounts, account)
	return nil
}

func NewInmemoryAccountRepository() *InMemoryAccountRepository {
	return &InMemoryAccountRepository{}
}
