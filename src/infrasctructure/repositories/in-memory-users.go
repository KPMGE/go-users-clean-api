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

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{}
}
