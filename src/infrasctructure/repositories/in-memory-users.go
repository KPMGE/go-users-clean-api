package repositories

import (
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
)

var users []*entities.User

type InMemoryUserRepository struct{}

func (repo *InMemoryUserRepository) Save(user *entities.User) error {
	users = append(users, user)
	return nil
}

func (repo *InMemoryUserRepository) CheckByEmail(email string) bool {
	for _, user := range users {
		if user.Email == email {
			return true
		}
	}
	return false
}

func (repo *InMemoryUserRepository) CheckByUserName(userName string) bool {
	for _, user := range users {
		if user.UserName == userName {
			return true
		}
	}
	return false
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{}
}
