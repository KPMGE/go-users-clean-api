package dto

type GetUserByIdUseCaseOutputDTO struct {
	ID       string
	Name     string
	Email    string
	UserName string
}

func NewGetUserByIdUseCaseOutputDTO(id string, name string, email string, userName string) *GetUserByIdUseCaseOutputDTO {
	return &GetUserByIdUseCaseOutputDTO{
		ID:       id,
		Email:    email,
		Name:     name,
		UserName: userName,
	}
}
