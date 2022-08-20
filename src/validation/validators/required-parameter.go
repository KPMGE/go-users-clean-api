package validators

import (
	"errors"
	"fmt"
	"reflect"
)

type RequiredParameterValidation struct {
	parameterName string
}

func getAttr(obj interface{}, fieldName string) reflect.Value {
	pointToStruct := reflect.ValueOf(obj) // addressable
	curStruct := pointToStruct.Elem()
	if curStruct.Kind() != reflect.Struct {
		panic("not a struct")
	}
	curField := curStruct.FieldByName(fieldName) // type: reflect.Value
	if !curField.IsValid() {
		panic("not found:" + fieldName)
	}
	return curField
}

func (r *RequiredParameterValidation) Validate(input any) error {
	val := getAttr(input, r.parameterName)
	if val.IsZero() {
		return errors.New(fmt.Sprintf("Missing field %s!", r.parameterName))
	}
	return nil
}

func NewRequiredParameterValidation(fieldName string) *RequiredParameterValidation {
	return &RequiredParameterValidation{
		parameterName: fieldName,
	}
}
