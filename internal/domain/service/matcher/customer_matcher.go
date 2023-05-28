package matcher

import (
	"context"
	"strings"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/event"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
)

type CustomerMatcherService struct {
}

func NewCustomerMatcherService() CustomerMatcherService {
	return CustomerMatcherService{}
}

func (s *CustomerMatcherService) MatchCustomerWithSegment(ctx context.Context, customer *model.Customer, segment model.Segment) (isMatch bool) {
	serializedCustomer := customer.Serialize()
	for _, condition := range segment.Conditions() {
		if !matchCondition(condition, serializedCustomer) {
			return false
		}
	}

	satisfiedSegment, err := model.NewCustomerSegment(
		model.NewCustomerSegmentInput{
			CustomerID:  customer.CustomerID(),
			WorkspaceID: customer.WorkspaceID(),
			SegmentID:   segment.SegmentID(),
		},
	)
	if err != nil {
		return false
	}
	customer.AppendCustomerSegment(*satisfiedSegment)

	customer.AggregateRoot.AppendEvent(event.CustomerMatched{
		DomainEventBase: domain.NewDomainEventBase(domain.NewDomainEventBaseInput{
			EventType:     event.CustomerEventOccurredEventType,
			OccurredAt:    satisfiedSegment.MatchedAt(),
			AggregateType: customer.AggregateType(),
			AggregateID:   customer.AggregateID(),
		}),
		CustomerID:  string(customer.CustomerID()),
		WorkspaceID: string(customer.WorkspaceID()),
		SegmentID:   string(satisfiedSegment.SegmentID()),
		MatchedAt:   satisfiedSegment.MatchedAt(),
	})

	return true
}

func matchCondition(condition model.Condition, customer model.SerializedCustomer) bool {

	type conditionHandler func(model.Condition, model.SerializedCustomer) bool
	mapMatchFunc := map[model.ConditionType]conditionHandler{
		model.SegmentConditionTypeContains:         isMatchByContainsCondition,
		model.SegmentConditionTypeEqual:            isMatchByEqualCondition,
		model.SegmentConditionTypeHasBeenPerformed: isMatchByHasBeenPerformedCondition,
		model.SegmentConditionTypeLessThan:         isMatchByLessThanCondition,
		model.SegmentConditionTypeMoreThan:         isMatchByMoreThanCondition,
	}

	return mapMatchFunc[condition.ConditionType](condition, customer)
}

func getCustomerAttributeValue(condition model.Condition, customer model.SerializedCustomer) (attributeValue any, find bool) {
	switch condition.Target {
	case model.ConditionTargetAttribute:
		customerAttributeValue, ok := customer.Attributes[condition.AttributeKey.String()]
		if !ok {
			return attributeValue, false
		}
		attributeValue = customerAttributeValue
	case model.ConditionTargetEvent:
		event, ok := customer.Events[condition.EventSlug]
		if !ok {
			return attributeValue, false
		}
		eventAttributeValue, ok := event[condition.AttributeKey.String()]
		if !ok {
			return attributeValue, false
		}
		attributeValue = eventAttributeValue
	}
	return attributeValue, true
}

func isMatchByContainsCondition(condition model.Condition, customer model.SerializedCustomer) (match bool) {
	attributeValue, ok := getCustomerAttributeValue(condition, customer)
	if !ok {
		return false
	}
	switch value := attributeValue.(type) {
	case string:
		match = strings.Contains(value, condition.AttributeValue.(string))
	default:
		return false
	}
	return match
}

func isMatchByEqualCondition(condition model.Condition, customer model.SerializedCustomer) (match bool) {
	attributeValue, ok := getCustomerAttributeValue(condition, customer)
	if !ok {
		return false
	}

	return attributeValue == condition.AttributeValue
}

func isMatchByHasBeenPerformedCondition(condition model.Condition, customer model.SerializedCustomer) (match bool) {
	_, match = customer.Events[condition.EventSlug]
	return match
}

func isMatchByLessThanCondition(condition model.Condition, customer model.SerializedCustomer) (match bool) {
	attributeValue, ok := getCustomerAttributeValue(condition, customer)
	if !ok {
		return false
	}
	switch value := attributeValue.(type) {
	case int:
		match = value < condition.AttributeValue.(int)
	case int16:
		match = value < condition.AttributeValue.(int16)
	case int32:
		match = value < condition.AttributeValue.(int32)
	case int64:
		match = value < condition.AttributeValue.(int64)
	case float32:
		match = value < condition.AttributeValue.(float32)
	case float64:
		match = value < condition.AttributeValue.(float64)
	default:
		return false
	}
	return match
}

func isMatchByMoreThanCondition(condition model.Condition, customer model.SerializedCustomer) (match bool) {
	attributeValue, ok := getCustomerAttributeValue(condition, customer)
	if !ok {
		return false
	}
	switch value := attributeValue.(type) {
	case int:
		match = value > condition.AttributeValue.(int)
	case int16:
		match = value > condition.AttributeValue.(int16)
	case int32:
		match = value > condition.AttributeValue.(int32)
	case int64:
		match = value > condition.AttributeValue.(int64)
	case float32:
		match = value > condition.AttributeValue.(float32)
	case float64:
		match = value > condition.AttributeValue.(float64)
	default:
		return false
	}
	return match
}
