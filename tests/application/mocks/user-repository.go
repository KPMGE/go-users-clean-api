package mocks_test

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type UserRepositorySpy struct {
	AddInput              *entities.User
	AddOutput             error
	CheckByEmailInput     string
	CheckByEmailOutput    bool
	CheckByUserNameInput  string
	CheckByUserNameOutput bool
}

func (repo *UserRepositorySpy) Save(user *entities.User) error {
	repo.AddInput = user
	return repo.AddOutput
}

func (repo *UserRepositorySpy) CheckByEmail(email string) bool {
	repo.CheckByEmailInput = email
	return repo.CheckByEmailOutput
}

func (repo *UserRepositorySpy) CheckByUserName(email string) bool {
	repo.CheckByUserNameInput = email
	return repo.CheckByUserNameOutput
}

func NewUserRepositorySpy() *UserRepositorySpy {
	return &UserRepositorySpy{}
}
