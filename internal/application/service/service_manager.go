package usecase

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/gateway"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
)

type ServiceManager struct {
	CrudCampaign  *CrudCampaignService
	CrudSegment   *CrudSegmentService
	CrudWorkspace *CrudWorkspaceService
	SaveActions   *SaveActionsService
}

func NewServiceManager(repoManager repository.RepositoryManager, gatewayManager gateway.GatewayManager) *ServiceManager {
	campaignRepo := repoManager.CampaignRepo()
	segmentRepo := repoManager.SegmentRepo()
	workspaceRepo := repoManager.WorkspaceRepo()

	return &ServiceManager{
		CrudCampaign:  NewCrudCampaignService(campaignRepo),
		CrudSegment:   NewCrudSegmentService(segmentRepo, campaignRepo),
		CrudWorkspace: NewCrudWorkspaceService(workspaceRepo),
		SaveActions:   NewSaveActionsService(campaignRepo),
	}
}
