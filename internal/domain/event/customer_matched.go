package event

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

var CustomerMatchedEventType = domain.EventType("customer-matched")

type CustomerMatched struct {
	*domain.DomainEventBase
	CustomerID  string
	WorkspaceID string
	SegmentID   string
	MatchedAt   time.Time
}
