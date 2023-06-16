package match_condition

import (
	"testing"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model/factory"
)

func (s *MatchConditionTestSuite) TestIsMatchByContainsCondition() {
	mockCustomer := s.customerFactory.CreateCustomer(factory.CreateCustomerInput{
		Name: "John Doe",
	})
	mockCustomer.AppendCustomerEvent("order-created", map[string]any{"product": "phone"})

	attributeNameContainsDoeCondition := s.conditionFactory.CreateContainsCondition(
		factory.CreateConditionInput{AttributeKey: "name", AttributeValue: "Doe"},
	)
	attributeNameContainsJaneCondition := s.conditionFactory.CreateContainsCondition(
		factory.CreateConditionInput{AttributeKey: "name", AttributeValue: "Jane"},
	)
	eventOrderAttributeProductContainsPhoneCondition := s.conditionFactory.CreateContainsCondition(
		factory.CreateConditionInput{Target: model.ConditionTargetEvent, EventSlug: "order-created", AttributeKey: "product", AttributeValue: "phone"},
	)
	eventOrderAttributeProductContainsMouseCondition := s.conditionFactory.CreateContainsCondition(
		factory.CreateConditionInput{Target: model.ConditionTargetEvent, EventSlug: "order-created", AttributeKey: "product", AttributeValue: "mouse"},
	)

	testCases := []testScenarios{
		{scenarioName: "test-match-condition-with-attribute-name", condition: attributeNameContainsDoeCondition, expectedMatched: true},
		{scenarioName: "test-not-match-condition-with-attribute-name", condition: attributeNameContainsJaneCondition, expectedMatched: false},
		{scenarioName: "test-match-condition-with-event-attribute-product", condition: eventOrderAttributeProductContainsPhoneCondition, expectedMatched: true},
		{scenarioName: "test-not-match-condition-with-event-attribute-product", condition: eventOrderAttributeProductContainsMouseCondition, expectedMatched: false},
	}

	for _, testCase := range testCases {
		s.Suite.T().Run(testCase.scenarioName, func(t *testing.T) {
			isMatch := isMatchByContainsCondition(testCase.condition, mockCustomer.Serialize())
			s.Suite.Equal(testCase.expectedMatched, isMatch)
		})
	}
}
