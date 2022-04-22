package dto

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type AddBookUseCaseInputDTO struct {
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	UserId      string  `json:"userId"`
}

type AddBookUseCaseOutputDTO struct {
	ID          string
	Title       string
	Author      string
	Price       float64
	Description string
	User        *entities.User
}
