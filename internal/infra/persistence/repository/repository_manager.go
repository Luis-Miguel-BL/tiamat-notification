package repository

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/config"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/logger"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/messaging"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/persistence/repository/implementation/dynamo"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/persistence/repository/implementation/in_memory"
)

type RepositoryManager struct {
	campaignRepo  repository.CampaignRepository
	customerRepo  repository.CustomerRepository
	journeyRepo   repository.JourneyRepository
	segmentRepo   repository.SegmentRepository
	workspaceRepo repository.WorkspaceRepository
}

func NewRepositoryManager(ctx context.Context, dispatcher messaging.AggregateEventDispatcher, cfg config.DBConfig, log logger.Logger) (*RepositoryManager, error) {
	if cfg.InMemory {
		return newInMemoryRepositoryManager(ctx, dispatcher, cfg, log)
	}
	dynamoClient, err := dynamo.NewDynamoClient(ctx, cfg, log)
	if err != nil {
		return nil, err
	}
	return &RepositoryManager{
		customerRepo: dynamo.NewDynamoCustomerRepo(dynamoClient, dispatcher),
	}, nil
}

func (r *RepositoryManager) CampaignRepo() repository.CampaignRepository {
	return r.campaignRepo
}
func (r *RepositoryManager) CustomerRepo() repository.CustomerRepository {
	return r.customerRepo
}
func (r *RepositoryManager) JourneyRepo() repository.JourneyRepository {
	return r.journeyRepo
}
func (r *RepositoryManager) SegmentRepo() repository.SegmentRepository {
	return r.segmentRepo
}
func (r *RepositoryManager) WorkspaceRepo() repository.WorkspaceRepository {
	return r.workspaceRepo
}

func newInMemoryRepositoryManager(ctx context.Context, dispatcher messaging.AggregateEventDispatcher, cfg config.DBConfig, log logger.Logger) (*RepositoryManager, error) {
	return &RepositoryManager{
		customerRepo: in_memory.NewInMemoryCustomerRepo(dispatcher),
	}, nil
}
