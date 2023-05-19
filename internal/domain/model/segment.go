package model

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

var AggregateTypeSegment = domain.AggregateType("segment")

type SegmentID string
type Segment struct {
	*domain.AggregateRoot
	segmentID   SegmentID
	workspaceID WorkspaceID
	slug        vo.Slug
	conditions  []Condition
	createdAt   time.Time
	updatedAt   time.Time
}

type NewSegmentInput struct {
	Slug        vo.Slug
	WorkspaceID WorkspaceID
	Conditions  []Condition
}

func NewSegment(input NewSegmentInput) (segment *Segment, err domain.DomainError) {
	segmentID := SegmentID(util.NewUUID())
	if input.WorkspaceID == "" {
		return segment, domain.NewInvalidEmptyParamError("WorkspaceID")
	}
	return &Segment{
		AggregateRoot: domain.NewAggregateRoot(AggregateTypeSegment, domain.AggregateID(segmentID)),
		segmentID:     segmentID,
		workspaceID:   input.WorkspaceID,
		slug:          input.Slug,
		conditions:    input.Conditions,
		createdAt:     time.Now(),
		updatedAt:     time.Now(),
	}, nil
}

func (e *Segment) SegmentID() SegmentID {
	return e.segmentID
}
func (e *Segment) Conditions() []Condition {
	return e.conditions
}
