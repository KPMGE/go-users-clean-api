package dto

type AddAccountOutputDTO struct {
	ID       string
	UserName string
	Email    string
}

type AddAccountInputDTO struct {
	UserName        string `json:"userName"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}
