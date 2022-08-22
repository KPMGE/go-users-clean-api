package protocols

import (
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
)

type ListAccountsRepository interface {
	ListAccounts() []entities.Account
}
