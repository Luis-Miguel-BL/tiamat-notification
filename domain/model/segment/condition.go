package segment

import "github.com/Luis-Miguel-BL/tiamat-notification/domain/vo"

type SerializedAttributes map[string]any
type SerializedCustomer struct {
	Attributes SerializedAttributes
	Events     map[vo.Slug]SerializedAttributes
}
type Condition interface {
	IsMatch(customer SerializedCustomer) (bool, error)
}

type ConditionType string
type ConditionTarget string

const (
	ConditionTargetEvent     ConditionTarget = "event"
	ConditionTargetAttribute ConditionTarget = "attribute"
)

type ConditionBase struct {
	ConditionType  ConditionType
	Target         ConditionTarget
	EventSlug      vo.Slug
	AttributeKey   vo.DotNotation
	AttributeValue any
}
