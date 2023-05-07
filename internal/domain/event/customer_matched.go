package event

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

var CustomerMatchedType = domain.EventType("customer-matched")

type CustomerMatched struct {
	*domain.DomainEventBase
	CustomerID  string
	WorkspaceID string
	SegmentID   string
	MatchedAt   time.Time
}
