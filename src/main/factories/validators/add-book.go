package validators

import (
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	"github.com/KPMGE/go-users-clean-api/src/validation/validators"
)

func MakeAddBookValidation() protocols.Validator {
	userIdField := validators.NewRequiredParameterValidation("UserId")
	titleField := validators.NewRequiredParameterValidation("Title")
	authorField := validators.NewRequiredParameterValidation("Author")
	descriptionField := validators.NewRequiredParameterValidation("Description")
	priceField := validators.NewRequiredParameterValidation("Price")
	composite := validators.NewValidationComposite([]protocols.Validator{
		titleField,
		userIdField,
		authorField,
		descriptionField,
		priceField,
	})
	return composite
}
