package match_condition

import (
	"strings"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
)

func isMatchByContainsCondition(condition model.Condition, customer model.SerializedCustomer) (match bool) {
	attributeValue, ok := getCustomerAttributeValue(condition, customer)
	if !ok {
		return false
	}
	value, ok := attributeValue.(string)
	if !ok {
		return false
	}
	match = strings.Contains(value, condition.AttributeValue.(string))
	return match
}
