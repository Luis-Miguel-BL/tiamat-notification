package domain

import "fmt"

type DomainError error

func NewInvalidEmptyParamError(paramName string) DomainError {
	return DomainError(fmt.Errorf("param %s cannot be empty", paramName))
}
func NewInvalidParamError(paramName string) DomainError {
	return DomainError(fmt.Errorf("param %s is invalid", paramName))
}
func NewNotFoundError(description string) DomainError {
	return DomainError(fmt.Errorf("%s not found", description))
}
func NewInvalidOperationError(operation string, description string) DomainError {
	return DomainError(fmt.Errorf("invalid operation %s: %s", operation, description))
}
