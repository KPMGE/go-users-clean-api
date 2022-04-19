package protocols

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type UserRepository interface {
	Save(user *entities.User) error
	CheckByEmail(email string) bool
	CheckByUserName(userName string) bool
	List() []*entities.User
	Delete(userId string) error
	CheckById(userId string) bool
	GetById(userId string) (*entities.User, error)
}
