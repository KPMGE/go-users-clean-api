package dto

type AddAccountOutputDTO struct {
	ID       string `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

type AddAccountInputDTO struct {
	UserName        string `json:"userName"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func NewAddAccountInputDTO(userName string, email string, password string, confirmPassword string) *AddAccountInputDTO {
	return &AddAccountInputDTO{
		UserName:        userName,
		Email:           email,
		Password:        password,
		ConfirmPassword: confirmPassword,
	}
}

func NewAddAccountOutputDTO(id string, userName string, email string) *AddAccountOutputDTO {
	return &AddAccountOutputDTO{
		ID:       id,
		UserName: userName,
		Email:    email,
	}
}
