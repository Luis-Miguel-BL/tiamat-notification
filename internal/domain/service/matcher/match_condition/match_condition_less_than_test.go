package match_condition

import (
	"testing"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model/factory"
)

func (s *MatchConditionTestSuite) TestIsMatchByLessThanCondition() {
	mockCustomer := s.customerFactory.CreateCustomer(factory.CreateCustomerInput{
		Name:             "John Doe",
		CustomAttributes: map[string]any{"age": 20},
	})
	mockCustomer.AppendCustomerEvent("order-created", map[string]any{"product": "phone", "price": 200.5})

	attributeAgeLessThan30Condition := s.conditionFactory.CreateLessThanCondition(
		factory.CreateConditionInput{AttributeKey: "age", AttributeValue: 30},
	)
	attributeAgeLessThan20Condition := s.conditionFactory.CreateLessThanCondition(
		factory.CreateConditionInput{AttributeKey: "age", AttributeValue: 20},
	)
	attributeAgeLessThan10Condition := s.conditionFactory.CreateLessThanCondition(
		factory.CreateConditionInput{AttributeKey: "age", AttributeValue: 10},
	)
	eventOrderPriceLessThan300Condition := s.conditionFactory.CreateLessThanCondition(
		factory.CreateConditionInput{EventSlug: "order-created", Target: model.ConditionTargetEvent, AttributeKey: "price", AttributeValue: 300.5},
	)
	eventOrderPriceLessThan100Condition := s.conditionFactory.CreateLessThanCondition(
		factory.CreateConditionInput{EventSlug: "order-created", Target: model.ConditionTargetEvent, AttributeKey: "price", AttributeValue: 100},
	)

	testCases := []testScenarios{
		{scenarioName: "test-match-condition-less-then-with-attribute-age", condition: attributeAgeLessThan30Condition, expectedMatched: true},
		{scenarioName: "test-match-condition-less-then-with-attribute-age", condition: attributeAgeLessThan20Condition, expectedMatched: false},
		{scenarioName: "test-match-condition-less-then-with-attribute-age", condition: attributeAgeLessThan10Condition, expectedMatched: false},
		{scenarioName: "test-match-condition-less-then-with-event-order", condition: eventOrderPriceLessThan300Condition, expectedMatched: true},
		{scenarioName: "test-match-condition-less-then-with-event-order", condition: eventOrderPriceLessThan100Condition, expectedMatched: false},
	}

	for _, testCase := range testCases {
		s.Suite.T().Run(testCase.scenarioName, func(t *testing.T) {
			isMatch := isMatchByLessThanCondition(testCase.condition, mockCustomer.Serialize())
			s.Suite.Equal(testCase.expectedMatched, isMatch)
		})
	}
}
