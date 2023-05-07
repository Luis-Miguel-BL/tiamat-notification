package vo

import "github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"

type CustomAttributes map[string]any

func NewCustomAttributes(customAttributes map[string]any) (CustomAttributes, domain.DomainError) {
	return CustomAttributes(customAttributes), nil
}

func (vo CustomAttributes) GetAttribute(attributeName string) (attributeValue any, find bool) {
	attributeValue, find = vo[attributeName]
	return attributeValue, find
}
