package protocols

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type UserRepository interface {
	Save(user *entities.User) error
	CheckByEmail(email string) bool
	CheckByUserName(userName string) bool
}
