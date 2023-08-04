package event

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

var CustomerMatchedEventType = domain.EventType("customer-matched")

type CustomerMatched struct {
	*domain.DomainEventBase
	CustomerID  string    `json:"customer_id"`
	WorkspaceID string    `json:"workspace_id"`
	SegmentID   string    `json:"segment_id"`
	MatchedAt   time.Time `json:"matched_at"`
}
