package fakedtos

import domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"

func MakeFakeLoginInputDTO() *domaindto.LoginInputDTO {
	return &domaindto.LoginInputDTO{
		UserName: "any_username",
		Email:    "any@email.com",
		Password: "any password",
	}
}
