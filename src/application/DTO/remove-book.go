package dto

type RemoveBookUseCaseOutputDTO struct {
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	UserId      string  `json:"userId"`
}
