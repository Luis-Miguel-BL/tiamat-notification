package match_condition

import (
	"testing"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model/factory"
	"github.com/stretchr/testify/suite"
)

type MatchConditionTestSuite struct {
	suite.Suite
	customerFactory  factory.CustomerFactory
	segmentFactory   factory.SegmentFactory
	conditionFactory factory.ConditionFactory
}

func (s *MatchConditionTestSuite) BeforeTest(suiteName, testName string) {
	s.customerFactory = factory.CustomerFactory{}
	s.segmentFactory = factory.SegmentFactory{}
	s.conditionFactory = factory.ConditionFactory{}
}
func TestMatchCondition(t *testing.T) {
	suite.Run(t, new(MatchConditionTestSuite))
}

type testScenarios struct {
	scenarioName    string
	condition       model.Condition
	expectedMatched bool
}
