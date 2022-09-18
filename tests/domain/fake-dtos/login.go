package fakedtos

import domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"

func MakeFakeLoginInputDTO() *domaindto.LoginInputDTO {
	return &domaindto.LoginInputDTO{
		Email:    "any@email.com",
		Password: "any password",
	}
}
