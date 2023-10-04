package usecase

import (
	ignition "github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/ignition"
	match "github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/match"
	workflow "github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/workflow"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/gateway"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
)

type UsecaseManager struct {
	CreateCustomerEvent *ignition.CreateCustomerEventUsecase
	SaveCustomer        *ignition.SaveCustomerUsecase
	MatchCustomer       *match.MatchCustomerUsecase
	ActionTracking      *workflow.ActionTrackingUsecase
	TriggerAction       *workflow.TriggerActionUsecase
}

func NewUsecaseManager(repoManager repository.RepositoryManager, gatewayManager gateway.GatewayManager) *UsecaseManager {
	campaignRepo := repoManager.CampaignRepo()
	customerRepo := repoManager.CustomerRepo()
	segmentRepo := repoManager.SegmentRepo()
	journeyRepo := repoManager.JourneyRepo()

	return &UsecaseManager{
		CreateCustomerEvent: ignition.NewCreateCustomerEventUsecase(customerRepo),
		SaveCustomer:        ignition.NewSaveCustomerUsecase(customerRepo),
		MatchCustomer:       match.NewMatchCustomerUsecase(customerRepo, segmentRepo, campaignRepo, journeyRepo),
		ActionTracking:      workflow.NewActionTrackingUsecase(journeyRepo, campaignRepo),
		TriggerAction:       workflow.NewTriggerActionUsecase(customerRepo, segmentRepo, campaignRepo, journeyRepo, gatewayManager),
	}
}
