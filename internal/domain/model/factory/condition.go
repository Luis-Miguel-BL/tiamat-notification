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
	if input.Target == "" {
		input.Target = model.ConditionTargetAttribute
	}
	var eventSlug vo.Slug
	if input.EventSlug != "" {
		eventSlug, _ = vo.NewSlug(input.EventSlug)
	}
	condition, err := model.NewCondition(model.NewConditionInput{
		ConditionTarget: input.Target,
		ConditionType:   input.ConditionType,
		EventSlug:       eventSlug,
		AttributeKey:    vo.DotNotation(input.AttributeKey),
		AttributeValue:  input.AttributeValue,
	})
	if err != nil {
		panic(err)
	}

	return condition
}

func (f *ConditionFactory) CreateContainsCondition(input CreateConditionInput) (condition model.Condition) {
	input.ConditionType = model.SegmentConditionTypeContains
	return f.CreateCondition(input)
}

func (f *ConditionFactory) CreateEqualCondition(input CreateConditionInput) (condition model.Condition) {
	input.ConditionType = model.SegmentConditionTypeEqual
	return f.CreateCondition(input)
}

func (f *ConditionFactory) CreateHasBeenPerformedCondition(input CreateConditionInput) (condition model.Condition) {
	input.ConditionType = model.SegmentConditionTypeHasBeenPerformed
	input.Target = model.ConditionTargetEvent
	return f.CreateCondition(input)
}
func (f *ConditionFactory) CreateLessThanCondition(input CreateConditionInput) (condition model.Condition) {
	input.ConditionType = model.SegmentConditionTypeLessThan
	return f.CreateCondition(input)
}
func (f *ConditionFactory) CreateMoreThanCondition(input CreateConditionInput) (condition model.Condition) {
	input.ConditionType = model.SegmentConditionTypeLessThan
	return f.CreateCondition(input)
}
