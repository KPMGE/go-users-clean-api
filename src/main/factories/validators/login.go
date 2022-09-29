package validators

import (
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	"github.com/KPMGE/go-users-clean-api/src/validation/validators"
)

func MakeLoginValidation() protocols.Validator {
	emailRequired := validators.NewRequiredParameterValidation("email")
	userNameRequired := validators.NewRequiredParameterValidation("userName")
	passwordRequired := validators.NewRequiredParameterValidation("password")

	composite := validators.NewValidationComposite([]protocols.Validator{
		emailRequired,
		userNameRequired,
		passwordRequired,
	})

	return composite
}
