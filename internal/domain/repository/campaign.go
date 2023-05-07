package repository

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
)

type CampaignRepository interface {
	Save(ctx context.Context, campaign model.Campaign) (err error)
	GetByID(ctx context.Context, campaignID model.CampaignID, workspaceID model.WorkspaceID) (campaign model.Campaign, err error)
	FindActiveCampaigns(ctx context.Context, workspaceID model.WorkspaceID) (campaigns []model.Campaign, err error)
}
