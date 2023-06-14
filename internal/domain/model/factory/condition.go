package factory

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

type ConditionFactory struct{}

type CreateConditionInput struct {
	ConditionType  model.ConditionType
	Target         model.ConditionTarget
	EventSlug      string
	AttributeKey   string
	AttributeValue any
}

func (f *ConditionFactory) CreateCondition(input CreateConditionInput) (condition model.Condition) {
	condition, _ = model.NewCondition(model.NewConditionInput{
		ConditionTarget: input.Target,
		ConditionType:   input.ConditionType,
		EventSlug:       vo.Slug(input.EventSlug),
		AttributeKey:    vo.DotNotation(input.AttributeKey),
		AttributeValue:  input.AttributeValue,
	})

	return condition
}

type CreateContainsConditionInput struct {
	Target         model.ConditionTarget
	EventSlug      string
	AttributeKey   string
	AttributeValue any
}

func (f *ConditionFactory) CreateContainsCondition(input CreateContainsConditionInput) (condition model.Condition) {
	if input.Target == "" {
		input.Target = model.ConditionTargetAttribute
	}
	condition, _ = model.NewCondition(model.NewConditionInput{
		ConditionTarget: input.Target,
		ConditionType:   model.SegmentConditionTypeContains,
		EventSlug:       vo.Slug(input.EventSlug),
		AttributeKey:    vo.DotNotation(input.AttributeKey),
		AttributeValue:  input.AttributeValue,
	})

	return condition
}
