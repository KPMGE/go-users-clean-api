package mocks_test

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type UserRepositorySpy struct {
	AddInput  *entities.User
	AddOutput error
}

func (repo *UserRepositorySpy) Save(user *entities.User) error {
	repo.AddInput = user
	return repo.AddOutput
}

func NewUserRepositorySpy() *UserRepositorySpy {
	return &UserRepositorySpy{}
}
