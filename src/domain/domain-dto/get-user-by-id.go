package domaindto

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type GetUserByIdUseCaseOutputDTO struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Email    string          `json:"email"`
	UserName string          `json:"userName"`
	Books    []entities.Book `json:"books"`
}

func NewGetUserByIdUseCaseOutputDTO(id string, name string, email string, userName string, books []entities.Book) *GetUserByIdUseCaseOutputDTO {
	return &GetUserByIdUseCaseOutputDTO{
		ID:       id,
		Email:    email,
		Name:     name,
		UserName: userName,
		Books:    books,
	}
}
