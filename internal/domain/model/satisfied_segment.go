package model

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type SatisfiedSegmentID string

func NewSatisfiedSegmentID(satisfiedSegmentID string) SatisfiedSegmentID {
	return SatisfiedSegmentID(satisfiedSegmentID)
}

type SatisfiedSegment struct {
	satisfiedSegmentID SatisfiedSegmentID
	segmentID          SegmentID
	workspaceID        WorkspaceID
	customerID         CustomerID
	matchedAt          time.Time
}

type NewSatisfiedSegmentInput struct {
	SegmentID   SegmentID
	CustomerID  CustomerID
	WorkspaceID WorkspaceID
}

func NewSatisfiedSegment(input NewSatisfiedSegmentInput) (satisfiedSegment *SatisfiedSegment, err error) {
	if input.SegmentID == "" {
		return satisfiedSegment, domain.NewInvalidEmptyParamError("SegmentID")
	}
	if input.WorkspaceID == "" {
		return satisfiedSegment, domain.NewInvalidEmptyParamError("WorkspaceID")
	}
	if input.CustomerID == "" {
		return satisfiedSegment, domain.NewInvalidEmptyParamError("CustomerID")
	}
	return &SatisfiedSegment{
		satisfiedSegmentID: SatisfiedSegmentID(util.NewUUID()),
		segmentID:          input.SegmentID,
		workspaceID:        input.WorkspaceID,
		customerID:         input.CustomerID,
		matchedAt:          time.Now(),
	}, nil
}

func (e *SatisfiedSegment) SegmentID() SegmentID {
	return e.segmentID
}
func (e *SatisfiedSegment) MatchedAt() time.Time {
	return e.matchedAt
}
