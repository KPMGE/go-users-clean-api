package protocols

type Validator interface {
	Validate(input any) error
}
