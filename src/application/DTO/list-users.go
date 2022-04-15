package dto

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type ListUsersDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"userNam"`
	Email    string `json:"email"`
}

func NewListUserDTO(id string, name string, userName string, email string) *ListUsersDTO {
	return &ListUsersDTO{
		ID:       id,
		Name:     name,
		UserName: userName,
		Email:    email,
	}
}

func MapListUsersDTO(users []*entities.User) []*ListUsersDTO {
	list := []*ListUsersDTO{}
	for _, user := range users {
		list = append(list, NewListUserDTO(user.ID, user.Name, user.UserName, user.Email))
	}
	return list
}
