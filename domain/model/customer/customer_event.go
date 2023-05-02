package customer

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/domain/vo"
)

type CustomerEventID string
type CustomerEvent struct {
	CustomerEventID CustomerEventID
	Slug            vo.Slug
	CustomData      map[string]interface{}
	CreatedAt       time.Time
}
