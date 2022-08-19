package domaindto

type AddUserInputDTO struct {
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

func NewAddUserInputDTO(name string, userName string, email string) *AddUserInputDTO {
	return &AddUserInputDTO{
		Name:     name,
		Email:    email,
		UserName: userName,
	}
}

type AddUserOutputDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

func NewAddUserOutputDTO(id string, name string, userName string, email string) *AddUserOutputDTO {
	return &AddUserOutputDTO{
		Name:     name,
		Email:    email,
		ID:       id,
		UserName: userName,
	}
}
