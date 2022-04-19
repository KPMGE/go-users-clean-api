package mocks_test

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type UserRepositorySpy struct {
	AddInput              *entities.User
	AddOutput             error
	CheckByEmailInput     string
	CheckByEmailOutput    bool
	CheckByUserNameInput  string
	CheckByUserNameOutput bool
	ListUsersOutput       []*entities.User
	DeleteInput           string
	DeleteOutput          error
	CheckByIdInput        string
	CheckByIdOuput        bool
	GetByidInput          string
	GetByidOutput         *entities.User
	GetByidError          error
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

func (repo *UserRepositorySpy) List() []*entities.User {
	return repo.ListUsersOutput
}

func (repo *UserRepositorySpy) Delete(userId string) error {
	repo.DeleteInput = userId
	return repo.DeleteOutput
}

func (repo *UserRepositorySpy) CheckById(userId string) bool {
	repo.CheckByIdInput = userId
	return repo.CheckByIdOuput
}

func (repo *UserRepositorySpy) GetById(userId string) (*entities.User, error) {
	repo.GetByidInput = userId
	return repo.GetByidOutput, repo.GetByidError
}

func NewUserRepositorySpy() *UserRepositorySpy {
	return &UserRepositorySpy{}
}
