package service

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
)

type TriggerService interface {
	TriggerCampaign(ctx context.Context, customer *model.Customer, campaign model.Campaign) (err error)
}

type triggerService struct {
	repo repository.CustomerRepository
}

func NewTriggerService(repo repository.CustomerRepository) TriggerService {
	return &triggerService{
		repo: repo,
	}
}

func (s *triggerService) TriggerCampaign(ctx context.Context, customer *model.Customer, campaign model.Campaign) (err error) {
	lastActionTriggered, alreadyTriggered := customer.GetLastActionTrigged(campaign.CampaignID())
	if alreadyTriggered {
		if !campaign.MustBeTriggered(lastActionTriggered.TriggeredAt()) {
			return
		}
	}

}
