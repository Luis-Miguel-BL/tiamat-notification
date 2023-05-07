package model

import (
	"time"
)

type CurrentSegmentID string

func NewCurrentSegmentID(currentSegmentID string) CurrentSegmentID {
	return CurrentSegmentID(currentSegmentID)
}

type CurrentSegment struct {
	currentSegmentID CurrentSegmentID
	segmentID        SegmentID
	matchedAt        time.Time
}

type NewCurrentSegmentInput struct {
	CurrentSegmentID CurrentSegmentID
	SegmentID        SegmentID
}

func NewCurrentSegment(input NewCurrentSegmentInput) *CurrentSegment {
	return &CurrentSegment{
		currentSegmentID: input.CurrentSegmentID,
		segmentID:        input.SegmentID,
		matchedAt:        time.Now(),
	}
}

func (e *CurrentSegment) SegmentID() SegmentID {
	return e.segmentID
}
func (e *CurrentSegment) MatchedAt() time.Time {
	return e.matchedAt
}
