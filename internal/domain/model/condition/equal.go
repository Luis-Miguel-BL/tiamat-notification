package condition

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
)

const SegmentConditionTypeEqual model.ConditionType = "equal"

type Equal struct {
	*model.ConditionBase
}

func (c *Equal) IsMatch(customer model.SerializedCustomer) (match bool) {
	attributeValue, ok := c.ConditionBase.GetCustomerAttributeValue(customer)
	if !ok {
		return false
	}

	return attributeValue == c.AttributeValue
}
