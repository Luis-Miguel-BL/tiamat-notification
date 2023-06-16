package match_condition

import (
	"testing"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model/factory"
)

func (s *MatchConditionTestSuite) TestIsMatchByEqualCondition() {
	mockCustomer := s.customerFactory.CreateCustomer(factory.CreateCustomerInput{
		Name: "John Doe",
		CustomAttributes: map[string]any{
			"age": 20,
			"phoneNumber": []map[string]interface{}{{
				"type":   "home",
				"number": "212 555-1234",
			}},
		},
	})
	mockCustomer.AppendCustomerEvent("order-created", map[string]any{"product": "phone", "price": 200.5})

	attributeNameEqualDoeCondition := s.conditionFactory.CreateEqualCondition(
		factory.CreateConditionInput{AttributeKey: "name", AttributeValue: "John Doe"},
	)
	attributeNameEqualJaneCondition := s.conditionFactory.CreateEqualCondition(
		factory.CreateConditionInput{AttributeKey: "name", AttributeValue: "Patrick Jane"},
	)
	attributeAgeEqual20Condition := s.conditionFactory.CreateEqualCondition(
		factory.CreateConditionInput{AttributeKey: "age", AttributeValue: 20},
	)
	attributeAgeEqual10Condition := s.conditionFactory.CreateEqualCondition(
		factory.CreateConditionInput{AttributeKey: "age", AttributeValue: 10},
	)
	attributePhoneNumberEqualHomeCondition := s.conditionFactory.CreateEqualCondition(
		factory.CreateConditionInput{AttributeKey: "phoneNumber[0].type", AttributeValue: "home"},
	)
	attributePhoneNumberEqualFaxCondition := s.conditionFactory.CreateEqualCondition(
		factory.CreateConditionInput{AttributeKey: "phoneNumber[0].type", AttributeValue: "fax"},
	)
	eventOrderAttributeProductEqualPhoneCondition := s.conditionFactory.CreateEqualCondition(
		factory.CreateConditionInput{Target: model.ConditionTargetEvent, EventSlug: "order-created", AttributeKey: "product", AttributeValue: "phone"},
	)
	eventOrderAttributeProductEqualMouseCondition := s.conditionFactory.CreateEqualCondition(
		factory.CreateConditionInput{Target: model.ConditionTargetEvent, EventSlug: "order-created", AttributeKey: "product", AttributeValue: "mouse"},
	)
	eventOrderAttributePriceEqual200Condition := s.conditionFactory.CreateEqualCondition(
		factory.CreateConditionInput{Target: model.ConditionTargetEvent, EventSlug: "order-created", AttributeKey: "price", AttributeValue: 200.5},
	)
	eventOrderAttributePriceEqual100Condition := s.conditionFactory.CreateEqualCondition(
		factory.CreateConditionInput{Target: model.ConditionTargetEvent, EventSlug: "order-created", AttributeKey: "price", AttributeValue: 100.5},
	)

	testCases := []testScenarios{
		{scenarioName: "test-match-condition-with-attribute-name", condition: attributeNameEqualDoeCondition, expectedMatched: true},
		{scenarioName: "test-not-match-condition-with-attribute-name", condition: attributeNameEqualJaneCondition, expectedMatched: false},
		{scenarioName: "test-match-condition-with-attribute-age", condition: attributeAgeEqual20Condition, expectedMatched: true},
		{scenarioName: "test-not-match-condition-with-attribute-age", condition: attributeAgeEqual10Condition, expectedMatched: false},
		{scenarioName: "test-match-condition-with-attribute-phone", condition: attributePhoneNumberEqualHomeCondition, expectedMatched: true},
		{scenarioName: "test-not-match-condition-with-attribute-phone", condition: attributePhoneNumberEqualFaxCondition, expectedMatched: false},
		{scenarioName: "test-match-condition-with-event-attribute-product", condition: eventOrderAttributeProductEqualPhoneCondition, expectedMatched: true},
		{scenarioName: "test-not-match-condition-with-event-attribute-product", condition: eventOrderAttributeProductEqualMouseCondition, expectedMatched: false},
		{scenarioName: "test-match-condition-with-event-attribute-price", condition: eventOrderAttributePriceEqual200Condition, expectedMatched: true},
		{scenarioName: "test-not-match-condition-with-event-attribute-price", condition: eventOrderAttributePriceEqual100Condition, expectedMatched: false},
	}

	for _, testCase := range testCases {
		s.Suite.T().Run(testCase.scenarioName, func(t *testing.T) {
			isMatch := isMatchByEqualCondition(testCase.condition, mockCustomer.Serialize())
			s.Suite.Equal(testCase.expectedMatched, isMatch)
		})
	}
}
