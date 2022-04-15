package mocks_test

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type FakeAccountRepository struct {
	Input                   *entities.Account
	CheckAccountOutput      bool
	CheckUserNameOutput     bool
	DeleteAccountByIdOutput bool
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

func (repo *FakeAccountRepository) DeleteAccountById(accountId string) bool {
	return repo.DeleteAccountByIdOutput
}

func NewFakeAccountRepository() *FakeAccountRepository {
	return &FakeAccountRepository{}
}
