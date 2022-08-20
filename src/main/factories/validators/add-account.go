package validators

import (
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	"github.com/KPMGE/go-users-clean-api/src/validation/validators"
)

func MakeAddAccountValidaton() protocols.Validator {
	userNameFiled := validators.NewRequiredParameterValidation("UserName")
	emailField := validators.NewRequiredParameterValidation("Email")
	passwordField := validators.NewRequiredParameterValidation("Password")
	confirmPasswordField := validators.NewRequiredParameterValidation("ConfirmPassword")
	composite := validators.NewValidationComposite([]protocols.Validator{userNameFiled, emailField, passwordField, confirmPasswordField})
	return composite
}
