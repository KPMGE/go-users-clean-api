package dto

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type GetUserByIdUseCaseOutputDTO struct {
	ID       string
	Name     string
	Email    string
	UserName string
	Books    []*entities.Book
}

func NewGetUserByIdUseCaseOutputDTO(id string, name string, email string, userName string, books []*entities.Book) *GetUserByIdUseCaseOutputDTO {
	return &GetUserByIdUseCaseOutputDTO{
		ID:       id,
		Email:    email,
		Name:     name,
		UserName: userName,
		Books:    books,
	}
}
