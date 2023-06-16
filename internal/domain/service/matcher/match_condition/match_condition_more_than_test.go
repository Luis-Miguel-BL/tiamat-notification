package match_condition

import (
	"testing"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model/factory"
)

func (s *MatchConditionTestSuite) TestIsMatchByMoreThanCondition() {
	mockCustomer := s.customerFactory.CreateCustomer(factory.CreateCustomerInput{
		Name:             "John Doe",
		CustomAttributes: map[string]any{"age": 20},
	})
	mockCustomer.AppendCustomerEvent("order-created", map[string]any{"product": "phone", "price": 200.5})

	attributeAgeMoreThan30Condition := s.conditionFactory.CreateMoreThanCondition(
		factory.CreateConditionInput{AttributeKey: "age", AttributeValue: 30},
	)
	attributeAgeMoreThan20Condition := s.conditionFactory.CreateMoreThanCondition(
		factory.CreateConditionInput{AttributeKey: "age", AttributeValue: 20},
	)
	attributeAgeMoreThan10Condition := s.conditionFactory.CreateMoreThanCondition(
		factory.CreateConditionInput{AttributeKey: "age", AttributeValue: 10},
	)
	eventOrderPriceMoreThan300Condition := s.conditionFactory.CreateMoreThanCondition(
		factory.CreateConditionInput{EventSlug: "order-created", Target: model.ConditionTargetEvent, AttributeKey: "price", AttributeValue: 300.5},
	)
	eventOrderPriceMoreThan100Condition := s.conditionFactory.CreateMoreThanCondition(
		factory.CreateConditionInput{EventSlug: "order-created", Target: model.ConditionTargetEvent, AttributeKey: "price", AttributeValue: 100},
	)

	testCases := []testScenarios{
		{scenarioName: "test-match-condition-more-then-with-attribute-age", condition: attributeAgeMoreThan30Condition, expectedMatched: false},
		{scenarioName: "test-match-condition-more-then-with-attribute-age", condition: attributeAgeMoreThan20Condition, expectedMatched: false},
		{scenarioName: "test-match-condition-more-then-with-attribute-age", condition: attributeAgeMoreThan10Condition, expectedMatched: true},
		{scenarioName: "test-match-condition-more-then-with-event-order", condition: eventOrderPriceMoreThan300Condition, expectedMatched: false},
		{scenarioName: "test-match-condition-more-then-with-event-order", condition: eventOrderPriceMoreThan100Condition, expectedMatched: true},
	}

	for _, testCase := range testCases {
		s.Suite.T().Run(testCase.scenarioName, func(t *testing.T) {
			isMatch := isMatchByMoreThanCondition(testCase.condition, mockCustomer.Serialize())
			s.Suite.Equal(testCase.expectedMatched, isMatch)
		})
	}
}
