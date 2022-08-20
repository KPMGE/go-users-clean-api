package validators

import (
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	"github.com/KPMGE/go-users-clean-api/src/validation/validators"
)

func MakeAddUserValidation() protocols.Validator {
	userNameFiled := validators.NewRequiredParameterValidation("UserName")
	nameFiled := validators.NewRequiredParameterValidation("Name")
	emailField := validators.NewRequiredParameterValidation("Email")
	composite := validators.NewValidationComposite([]protocols.Validator{userNameFiled, emailField, nameFiled })
	return composite
}
