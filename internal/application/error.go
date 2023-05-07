package application

import "fmt"

type ApplicationError error

func NewInvalidEmptyParamError(paramName string) ApplicationError {
	return ApplicationError(fmt.Errorf("param %s cannot be empty", paramName))
}
