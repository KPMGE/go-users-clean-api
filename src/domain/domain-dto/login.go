package domaindto

type LoginInputDTO struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
