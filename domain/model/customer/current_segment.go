package customer

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/domain/model/segment"
)

type CurrentSegmentID string

type CurrentSegment struct {
	CurrentSegmentID CurrentSegmentID
	SegmentID        segment.SegmentID
	MatchedAt        time.Time
}
