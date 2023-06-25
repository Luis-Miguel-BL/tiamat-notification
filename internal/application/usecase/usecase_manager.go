package usecase

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/repository"
	control "github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/control"
	ignition "github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/ignition"
	match "github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/match"
	workflow "github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/workflow"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/gateway"
)

type UsecaseManager struct {
	CrudCampaign        *control.CrudCampaignUsecase
	CrudSegment         *control.CrudSegmentUsecase
	CrudWorkspace       *control.CrudWorkspaceUsecase
	SaveActions         *control.SaveActionsUsecase
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
	workspaceRepo := repoManager.WorkspaceRepo()

	return &UsecaseManager{
		CrudCampaign:        control.NewCrudCampaignUsecase(campaignRepo),
		CrudSegment:         control.NewCrudSegmentUsecase(segmentRepo, campaignRepo),
		CrudWorkspace:       control.NewCrudWorkspaceUsecase(workspaceRepo),
		SaveActions:         control.NewSaveActionsUsecase(campaignRepo),
		CreateCustomerEvent: ignition.NewCreateCustomerEventUsecase(customerRepo),
		SaveCustomer:        ignition.NewSaveCustomerUsecase(customerRepo),
		MatchCustomer:       match.NewMatchCustomerUsecase(customerRepo, segmentRepo, campaignRepo, journeyRepo),
		ActionTracking:      workflow.NewActionTrackingUsecase(journeyRepo, campaignRepo),
		TriggerAction:       workflow.NewTriggerActionUsecase(customerRepo, segmentRepo, campaignRepo, journeyRepo, gatewayManager),
	}
}
