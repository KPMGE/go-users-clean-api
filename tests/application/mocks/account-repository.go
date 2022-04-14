package mocks_test

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type FakeAccountRepository struct {
	Input                 *entities.Account
	CheckAccountOutput    bool
	CheckUserNameOutput   bool
	FindAccountByIdOutput *entities.Account
}

func (repo *FakeAccountRepository) CheckAccountByEmail(email string) bool {
	return repo.CheckAccountOutput
}

func (repo *FakeAccountRepository) CheckAccountByUserName(userName string) bool {
	return repo.CheckUserNameOutput
}

func (repo *FakeAccountRepository) Save(account *entities.Account) error {
	repo.Input = account
	return nil
}

func (repo *FakeAccountRepository) FindAccountById(accountId string) *entities.Account {
	return repo.FindAccountByIdOutput
}

func NewFakeAccountRepository() *FakeAccountRepository {
	return &FakeAccountRepository{}
}
