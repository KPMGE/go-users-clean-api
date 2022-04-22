package dto

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type AddBookUseCaseInputDTO struct {
	Title       string
	Author      string
	Price       float64
	Description string
	UserId      string
}

type AddBookUseCaseOutputDTO struct {
	ID          string
	Title       string
	Author      string
	Price       float64
	Description string
	User        *entities.User
}

func NewAddBookUseCaseInputDTO(title string, author string, price float64, description string, userId string) *AddBookUseCaseInputDTO {
	return &AddBookUseCaseInputDTO{
		Title:       title,
		Author:      author,
		Price:       price,
		Description: description,
		UserId:      userId,
	}
}
