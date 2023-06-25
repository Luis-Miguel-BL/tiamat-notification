package errors

import "fmt"

func NewInvalidEmptyParamError(paramName string) error {
	return fmt.Errorf("param %s cannot be empty", paramName)
}

func NewInvalidParamError(paramName string) error {
	return fmt.Errorf("param %s is invalid", paramName)
}
