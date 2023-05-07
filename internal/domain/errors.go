package domain

import "fmt"

type DomainError error

func NewInvalidEmptyParamError(paramName string) DomainError {
	return DomainError(fmt.Errorf("param %s cannot be empty", paramName))
}
func NewInvalidParamError(paramName string) DomainError {
	return DomainError(fmt.Errorf("param %s is invalid", paramName))
}
