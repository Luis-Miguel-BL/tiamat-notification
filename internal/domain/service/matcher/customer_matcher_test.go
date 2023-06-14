package matcher

import (
	"testing"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model/factory"
	"github.com/stretchr/testify/suite"
)

type CustomerMatcherServiceTestSuite struct {
	suite.Suite
	svc              CustomerMatcherService
	customerFactory  factory.CustomerFactory
	segmentFactory   factory.SegmentFactory
	conditionFactory factory.ConditionFactory
}

func (s *CustomerMatcherServiceTestSuite) BeforeTest(suiteName, testName string) {
	s.svc = NewCustomerMatcherService()
	s.customerFactory = factory.CustomerFactory{}
	s.segmentFactory = factory.SegmentFactory{}
	s.conditionFactory = factory.ConditionFactory{}
}
func TestCustomerMatcherService(t *testing.T) {
	suite.Run(t, new(CustomerMatcherServiceTestSuite))
}

func (s *CustomerMatcherServiceTestSuite) TestIsMatchByContainsCondition() {
	mockCustomer := s.customerFactory.CreateCustomer(factory.CreateCustomerInput{
		Name: "John Doe",
	})
	mockCustomer.AppendCustomerEvent("order-created", map[string]any{"product": "phone"})

	attributeNameContainsDoeCondition := s.conditionFactory.CreateContainsCondition(
		factory.CreateContainsConditionInput{AttributeKey: "name", AttributeValue: "Doe"},
	)
	attributeNameContainsJaneCondition := s.conditionFactory.CreateContainsCondition(
		factory.CreateContainsConditionInput{AttributeKey: "name", AttributeValue: "Jane"},
	)
	eventOrderAttributeProductContainsPhoneCondition := s.conditionFactory.CreateContainsCondition(
		factory.CreateContainsConditionInput{Target: model.ConditionTargetEvent, EventSlug: "order-created", AttributeKey: "product", AttributeValue: "phone"},
	)
	eventOrderAttributeProductContainsMouseCondition := s.conditionFactory.CreateContainsCondition(
		factory.CreateContainsConditionInput{Target: model.ConditionTargetEvent, EventSlug: "order-created", AttributeKey: "product", AttributeValue: "mouse"},
	)

	type testScenarios struct {
		scenarioName    string
		condition       model.Condition
		expectedMatched bool
	}
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
