package match_condition

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
)

func isMatchByEqualCondition(condition model.Condition, customer model.SerializedCustomer) (match bool) {
	attributeValue, ok := getCustomerAttributeValue(condition, customer)
	if !ok {
		return false
	}

	return attributeValue == condition.AttributeValue
}
