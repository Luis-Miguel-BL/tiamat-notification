package model

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

var AggregateTypeCampaign = domain.AggregateType("campaign")

type CampaignID string
type Campaign struct {
	*domain.AggregateRoot
	campaignID         CampaignID
	slug               vo.Slug
	sendToUnsubscribed bool
	isActive           bool
	firstActionID      ActionID
	actions            map[ActionID]Action
	triggers           []SegmentID
	filters            []SegmentID
	createdAt          time.Time
	updatedAt          time.Time
}

func (e *Campaign) Triggers() []SegmentID {
	return e.triggers
}
func (e *Campaign) Filters() []SegmentID {
	return e.filters
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
