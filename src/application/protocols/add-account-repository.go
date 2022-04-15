package protocols

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type AccountRepository interface {
	CheckAccountByEmail(email string) bool
	CheckAccountByUserName(userName string) bool
	Save(account *entities.Account) error
	DeleteAccountById(accountId string) bool
}
