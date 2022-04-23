package dto

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type ListUsersDTO struct {
	ID       string           `json:"id"`
	Name     string           `json:"name"`
	UserName string           `json:"userNam"`
	Email    string           `json:"email"`
	Books    []*entities.Book `json:"books"`
}

func NewListUserDTO(id string, name string, userName string, email string, books []*entities.Book) *ListUsersDTO {
	return &ListUsersDTO{
		ID:       id,
		Name:     name,
		UserName: userName,
		Email:    email,
		Books:    books,
	}
}

func MapListUsersDTO(users []*entities.User) []*ListUsersDTO {
	list := []*ListUsersDTO{}
	for _, user := range users {
		list = append(list, NewListUserDTO(user.ID, user.Name, user.UserName, user.Email, user.Books))
	}
	return list
}
