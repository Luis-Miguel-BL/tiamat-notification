package repository

type RepositoryManager interface {
	CampaignRepo() CampaignRepository
	CustomerRepo() CustomerRepository
	JourneyRepo() JourneyRepository
	SegmentRepo() SegmentRepository
	WorkspaceRepo() WorkspaceRepository
}
