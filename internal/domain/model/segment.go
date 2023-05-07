package model

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

var AggregateTypeSegment = domain.AggregateType("segment")

type SegmentID string
type Segment struct {
	*domain.AggregateRoot
	SegmentID  SegmentID
	Slug       vo.Slug
	Conditions []Condition
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
