package condition

import "github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"

const SegmentConditionTypeMoreThan model.ConditionType = "contains"

type MoreThan struct {
	*model.ConditionBase
}

func (c *MoreThan) IsMatch(customer model.SerializedCustomer) (match bool) {
	attributeValue, ok := c.ConditionBase.GetCustomerAttributeValue(customer)
	if !ok {
		return false
	}
	switch value := attributeValue.(type) {
	case int:
		match = value > c.AttributeValue.(int)
	case int16:
		match = value > c.AttributeValue.(int16)
	case int32:
		match = value > c.AttributeValue.(int32)
	case int64:
		match = value > c.AttributeValue.(int64)
	case float32:
		match = value > c.AttributeValue.(float32)
	case float64:
		match = value > c.AttributeValue.(float64)
	default:
		return false
	}
	return match
}
