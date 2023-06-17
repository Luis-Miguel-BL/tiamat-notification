package model

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

type CustomerSegment struct {
	segmentID   SegmentID
	workspaceID WorkspaceID
	customerID  CustomerID
	matchedAt   time.Time
}

type NewCustomerSegmentInput struct {
	SegmentID   SegmentID
	CustomerID  CustomerID
	WorkspaceID WorkspaceID
}

func NewCustomerSegment(input NewCustomerSegmentInput) (satisfiedSegment *CustomerSegment, err error) {
	if input.SegmentID == "" {
		return satisfiedSegment, domain.NewInvalidEmptyParamError("SegmentID")
	}
	if input.WorkspaceID == "" {
		return satisfiedSegment, domain.NewInvalidEmptyParamError("WorkspaceID")
	}
	if input.CustomerID == "" {
		return satisfiedSegment, domain.NewInvalidEmptyParamError("CustomerID")
	}
	return &CustomerSegment{
		segmentID:   input.SegmentID,
		workspaceID: input.WorkspaceID,
		customerID:  input.CustomerID,
		matchedAt:   time.Now(),
	}, nil
}

func (e *CustomerSegment) SegmentID() SegmentID {
	return e.segmentID
}
func (e *CustomerSegment) CustomerID() CustomerID {
	return e.customerID
}
func (e *CustomerSegment) WorkspaceID() WorkspaceID {
	return e.workspaceID
}
func (e *CustomerSegment) MatchedAt() time.Time {
	return e.matchedAt
}
