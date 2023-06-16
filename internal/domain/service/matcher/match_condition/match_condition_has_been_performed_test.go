package match_condition

import (
	"testing"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model/factory"
)

func (s *MatchConditionTestSuite) TestIsMatchByHasBeenPerformedCondition() {
	mockCustomer := s.customerFactory.CreateCustomer(factory.CreateCustomerInput{
		Name:             "John Doe",
		CustomAttributes: map[string]any{"age": 20},
	})
	mockCustomer.AppendCustomerEvent("order-created", map[string]any{"product": "phone", "price": 200.5})

	eventOrderHasBeenPerformedCondition := s.conditionFactory.CreateHasBeenPerformedCondition(
		factory.CreateConditionInput{EventSlug: "order-created"},
	)

	eventDeliveryHasBeenPerformedCondition := s.conditionFactory.CreateHasBeenPerformedCondition(
		factory.CreateConditionInput{EventSlug: "delivery-created"},
	)

	testCases := []testScenarios{
		{scenarioName: "test-match-condition-with-event-order", condition: eventOrderHasBeenPerformedCondition, expectedMatched: true},
		{scenarioName: "test-not-match-condition-with-event-delivery", condition: eventDeliveryHasBeenPerformedCondition, expectedMatched: false},
	}

	for _, testCase := range testCases {
		s.Suite.T().Run(testCase.scenarioName, func(t *testing.T) {
			isMatch := isMatchByHasBeenPerformedCondition(testCase.condition, mockCustomer.Serialize())
			s.Suite.Equal(testCase.expectedMatched, isMatch)
		})
	}
}
