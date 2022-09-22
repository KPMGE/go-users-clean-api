package presentationerrors

import "errors"

func NewBlankBodyError() error {
	return errors.New("json body is required, but a blank one was provided!")
}
