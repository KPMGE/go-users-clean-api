package dto

type AddUserInputDTO struct {
	Name     string
	UserName string
	Email    string
}

func NewAddUserInputDTO(name string, userName string, email string) *AddUserInputDTO {
	return &AddUserInputDTO{
		Name:     name,
		Email:    email,
		UserName: userName,
	}
}

type AddUserOutputDTO struct {
	ID       string
	Name     string
	UserName string
	Email    string
}

func NewAddUserOutputDTO(id string, name string, userName string, email string) *AddUserOutputDTO {
	return &AddUserOutputDTO{
		Name:     name,
		Email:    email,
		ID:       id,
		UserName: userName,
	}
}
