package usecases

type DeleteUserUseCase interface {
	DeleteUser(userId string) (string, error)
}
