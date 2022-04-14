package protocols

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type UserRepository interface {
	Save(user *entities.User) error
}
