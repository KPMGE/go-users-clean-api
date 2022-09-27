package controllers

import (
	"encoding/json"
	"errors"

	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	usecases "github.com/KPMGE/go-users-clean-api/src/domain/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	presentationerrors "github.com/KPMGE/go-users-clean-api/src/presentation/presentation-errors"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

type LoginController struct {
	srv       usecases.LoginUseCase
	validator protocols.Validator
}

func (c *LoginController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	if request == nil || request.Body == nil {
		return helpers.BadRequest(presentationerrors.NewBlankBodyError())
	}

	var input domaindto.LoginInputDTO
	err := json.Unmarshal(request.Body, &input)

	if err != nil {
		return helpers.ServerError(errors.New("error when parsing json body!"))
	}

	err = c.validator.Validate(&input)
	if err != nil {
		return helpers.BadRequest(err)
	}

	token, err := c.srv.Login(&input)

	if err != nil {
		return helpers.ServerError(err)
	}

	return helpers.Ok(token)
}

func NewLoginController(srv usecases.LoginUseCase, validator protocols.Validator) *LoginController {
	return &LoginController{
		srv:       srv,
		validator: validator,
	}
}
