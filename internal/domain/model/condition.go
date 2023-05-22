package model

import (
	"reflect"
	"strings"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

type ConditionHandler func(SerializedCustomer) bool

type ConditionType string

const (
	SegmentConditionTypeContains         ConditionType = "contains"
	SegmentConditionTypeEqual            ConditionType = "equal"
	SegmentConditionTypeHasBeenPerformed ConditionType = "hass-been-performed"
	SegmentConditionTypeLessThan         ConditionType = "less-than"
	SegmentConditionTypeMoreThan         ConditionType = "more-than"
)

// var AvailableKind = "sfs"
var ConditionAvailableAttrKind = map[ConditionType][]reflect.Kind{
	SegmentConditionTypeContains:         {reflect.String},
	SegmentConditionTypeEqual:            {reflect.String, reflect.Bool, reflect.Int, reflect.Float64},
	SegmentConditionTypeHasBeenPerformed: {},
	SegmentConditionTypeLessThan:         {reflect.Int, reflect.Int64, reflect.Float32, reflect.Float64},
	SegmentConditionTypeMoreThan:         {reflect.Int, reflect.Int64, reflect.Float32, reflect.Float64},
}

type ConditionTarget string

const (
	ConditionTargetEvent     ConditionTarget = "event"
	ConditionTargetAttribute ConditionTarget = "attribute"
)

type Condition struct {
	ConditionType  ConditionType
	Target         ConditionTarget
	EventSlug      vo.Slug
	AttributeKey   vo.DotNotation
	AttributeValue any
}
type NewConditionInput struct {
	ConditionTarget ConditionTarget
	ConditionType   ConditionType
	EventSlug       vo.Slug
	AttributeKey    vo.DotNotation
	AttributeValue  any
}

func NewCondition(input NewConditionInput) (condition Condition, err domain.DomainError) {
	switch input.ConditionTarget {
	case ConditionTargetAttribute:
		condition.Target = ConditionTargetAttribute
	case ConditionTargetEvent:
		condition.Target = ConditionTargetEvent
	default:
		return condition, domain.NewInvalidParamError("condition_target")
	}

	switch input.ConditionType {
	case SegmentConditionTypeContains:
		condition.ConditionType = SegmentConditionTypeContains
	case SegmentConditionTypeEqual:
		condition.ConditionType = SegmentConditionTypeEqual
	case SegmentConditionTypeHasBeenPerformed:
		condition.ConditionType = SegmentConditionTypeHasBeenPerformed
	case SegmentConditionTypeLessThan:
		condition.ConditionType = SegmentConditionTypeLessThan
	case SegmentConditionTypeMoreThan:
		condition.ConditionType = SegmentConditionTypeMoreThan
	default:
		return condition, domain.NewInvalidParamError("condition_type")
	}

	isMissingEventSlug := condition.Target == ConditionTargetEvent && input.EventSlug == ""
	if isMissingEventSlug {
		return condition, domain.NewInvalidEmptyParamError("event_slug")
	}

	isMissingAttrKey := condition.ConditionType != SegmentConditionTypeHasBeenPerformed && input.AttributeKey == ""
	if isMissingAttrKey {
		return condition, domain.NewInvalidEmptyParamError("attribute_key")
	}

	availablesAttrValueKind := ConditionAvailableAttrKind[condition.ConditionType]
	isAvailableAttrKind := false

	attrValueKind := reflect.TypeOf(input.AttributeValue).Kind()
	for _, availableKind := range availablesAttrValueKind {
		if attrValueKind == availableKind {
			isAvailableAttrKind = true
		}
	}
	if !isAvailableAttrKind {
		return condition, domain.NewInvalidParamError("attribute_value")
	}
	condition.EventSlug = input.EventSlug
	condition.AttributeKey = input.AttributeKey
	condition.AttributeValue = input.AttributeValue

	return condition, nil
}

func (e *Condition) IsMatch(customer SerializedCustomer) bool {
	mapMatchFunc := map[ConditionType]ConditionHandler{
		SegmentConditionTypeContains:         e.isMatchByContainsCondition,
		SegmentConditionTypeEqual:            e.isMatchByEqualCondition,
		SegmentConditionTypeHasBeenPerformed: e.isMatchByHasBeenPerformedCondition,
		SegmentConditionTypeLessThan:         e.isMatchByLessThanCondition,
		SegmentConditionTypeMoreThan:         e.isMatchByMoreThanCondition,
	}

	return mapMatchFunc[e.ConditionType](customer)
}

func (e *Condition) getCustomerAttributeValue(customer SerializedCustomer) (attributeValue any, find bool) {
	switch e.Target {
	case ConditionTargetAttribute:
		customerAttributeValue, ok := customer.Attributes[e.AttributeKey.String()]
		if !ok {
			return attributeValue, false
		}
		attributeValue = customerAttributeValue
	case ConditionTargetEvent:
		event, ok := customer.Events[e.EventSlug]
		if !ok {
			return attributeValue, false
		}
		eventAttributeValue, ok := event[e.AttributeKey.String()]
		if !ok {
			return attributeValue, false
		}
		attributeValue = eventAttributeValue
	}
	return attributeValue, true
}

func (e *Condition) isMatchByContainsCondition(customer SerializedCustomer) (match bool) {
	attributeValue, ok := e.getCustomerAttributeValue(customer)
	if !ok {
		return false
	}
	switch value := attributeValue.(type) {
	case string:
		match = strings.Contains(value, e.AttributeValue.(string))
	default:
		return false
	}
	return match
}

func (e *Condition) isMatchByEqualCondition(customer SerializedCustomer) (match bool) {
	attributeValue, ok := e.getCustomerAttributeValue(customer)
	if !ok {
		return false
	}

	return attributeValue == e.AttributeValue
}

func (e *Condition) isMatchByHasBeenPerformedCondition(customer SerializedCustomer) (match bool) {
	_, match = customer.Events[e.EventSlug]
	return match
}

func (e *Condition) isMatchByLessThanCondition(customer SerializedCustomer) (match bool) {
	attributeValue, ok := e.getCustomerAttributeValue(customer)
	if !ok {
		return false
	}
	switch value := attributeValue.(type) {
	case int:
		match = value < e.AttributeValue.(int)
	case int16:
		match = value < e.AttributeValue.(int16)
	case int32:
		match = value < e.AttributeValue.(int32)
	case int64:
		match = value < e.AttributeValue.(int64)
	case float32:
		match = value < e.AttributeValue.(float32)
	case float64:
		match = value < e.AttributeValue.(float64)
	default:
		return false
	}
	return match
}

func (e *Condition) isMatchByMoreThanCondition(customer SerializedCustomer) (match bool) {
	attributeValue, ok := e.getCustomerAttributeValue(customer)
	if !ok {
		return false
	}
	switch value := attributeValue.(type) {
	case int:
		match = value > e.AttributeValue.(int)
	case int16:
		match = value > e.AttributeValue.(int16)
	case int32:
		match = value > e.AttributeValue.(int32)
	case int64:
		match = value > e.AttributeValue.(int64)
	case float32:
		match = value > e.AttributeValue.(float32)
	case float64:
		match = value > e.AttributeValue.(float64)
	default:
		return false
	}
	return match
}
