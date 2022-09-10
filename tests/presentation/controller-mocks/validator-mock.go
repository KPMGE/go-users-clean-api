package controllermocks_test

type ValidatorMock struct {
	Output error
}

func (v *ValidatorMock) Validate(input any) error {
	return v.Output
}
