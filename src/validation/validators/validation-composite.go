package validators

import "github.com/KPMGE/go-users-clean-api/src/presentation/protocols"

type ValidationComposite struct {
	validators []protocols.Validator
}

func (v *ValidationComposite) Validate(input any) error {
	for _, validator := range v.validators {
		err := validator.Validate(input)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewValidationComposite(validators []protocols.Validator) *ValidationComposite {
	return &ValidationComposite{
		validators: validators,
	}
}
