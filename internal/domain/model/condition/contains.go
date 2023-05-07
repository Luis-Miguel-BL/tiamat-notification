package condition

import (
	"strings"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
)

const SegmentConditionTypeContains model.ConditionType = "contains"

type Contains struct {
	*model.ConditionBase
}

func (c *Contains) IsMatch(customer model.SerializedCustomer) (match bool) {
	attributeValue, ok := c.ConditionBase.GetCustomerAttributeValue(customer)
	if !ok {
		return false
	}
	switch value := attributeValue.(type) {
	case string:
		match = strings.Contains(value, c.AttributeValue.(string))
	default:
		return false
	}
	return match
}
