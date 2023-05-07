package model

import "github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"

type Condition interface {
	IsMatch(customer SerializedCustomer) bool
}

type ConditionType string
type ConditionTarget string

const (
	ConditionTargetEvent     ConditionTarget = "event"
	ConditionTargetAttribute ConditionTarget = "attribute"
)

type ConditionBase struct {
	ConditionType  ConditionType
	Target         ConditionTarget
	EventSlug      vo.Slug
	AttributeKey   vo.DotNotation
	AttributeValue any
}

func (c *ConditionBase) GetCustomerAttributeValue(customer SerializedCustomer) (attributeValue any, find bool) {
	switch c.Target {
	case ConditionTargetAttribute:
		customerAttributeValue, ok := customer.Attributes[c.AttributeKey.String()]
		if !ok {
			return attributeValue, false
		}
		attributeValue = customerAttributeValue
	case ConditionTargetEvent:
		event, ok := customer.Events[c.EventSlug]
		if !ok {
			return attributeValue, false
		}
		eventAttributeValue, ok := event[c.AttributeKey.String()]
		if !ok {
			return attributeValue, false
		}
		attributeValue = eventAttributeValue
	}
	return attributeValue, true
}
