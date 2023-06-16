package match_condition

import "github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"

func Match(condition model.Condition, customer model.SerializedCustomer) bool {

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
