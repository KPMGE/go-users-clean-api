package mocks_test

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type UserRepositorySpy struct {
	AddInput          *entities.User
	AddOutput         error
	FindByEmailInput  string
	FindByEmailOutput bool
}

func (repo *UserRepositorySpy) Save(user *entities.User) error {
	repo.AddInput = user
	return repo.AddOutput
}

func (repo *UserRepositorySpy) FindByEmail(email string) bool {
	repo.FindByEmailInput = email
	return repo.FindByEmailOutput
}

func NewUserRepositorySpy() *UserRepositorySpy {
	return &UserRepositorySpy{}
}
