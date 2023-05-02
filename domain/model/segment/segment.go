package segment

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/domain/vo"
)

var AggregateType = domain.AggregateType("segment")

type SegmentID string
type Segment struct {
	*domain.Aggregate
	SegmentID  SegmentID
	Slug       vo.Slug
	Conditions []Condition
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
