package protocols

import domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"

type ListAccountsRepository interface {
	ListAccounts() []domaindto.ListAccountsOutputDTO
}
