package usecases

type RemoveAccountUseCase interface {
	RemoveAccount(accountId string) (string, error)
}
