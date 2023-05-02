package campaign

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/domain/vo"
)

var AggregateType = domain.AggregateType("campaign")

type CampaignID string
type Campaign struct {
	*domain.Aggregate
	CampaignID         CampaignID
	Slug               vo.Slug
	SendToUnsubscribed bool
	FirstActionID      ActionID
	Actions            map[ActionID]Action
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

// func NewCampaign(workspaceID string, Slug string) (customer *Campaign) {
// 	return &Campaign{
// 		CampaignID: vo.ID(workspaceID),
// 		Slug:      vo.Slug(Slug),
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}
// }

// func (e *Campaign) Validate() error {
// 	if err := e.CampaignID.Validate(); err != nil {
// 		return err
// 	}
// 	if err := e.Slug.Validate(); err != nil {
// 		return err
// 	}
// 	return nil
// }
