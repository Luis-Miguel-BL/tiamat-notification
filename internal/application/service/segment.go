package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/service/input"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type CrudSegmentService struct {
	segmentRepo  repository.SegmentRepository
	campaignRepo repository.CampaignRepository
}

func NewCrudSegmentService(segmentRepo repository.SegmentRepository, campaignRepo repository.CampaignRepository) *CrudSegmentService {
	return &CrudSegmentService{
		segmentRepo:  segmentRepo,
		campaignRepo: campaignRepo,
	}
}

func (uc *CrudSegmentService) CreateSegment(ctx context.Context, input input.CreateSegmentInput) (err error) {
	err = input.Validate()
	if err != nil {
		return err
	}

	segmentSlug, err := vo.NewSlug(input.Slug)
	if err != nil {
		return err
	}
	workspaceID := model.WorkspaceID(input.WorkspaceID)

	segmentConditions, err := parseConditions(input.Conditions)
	if err != nil {
		return err
	}

	segmentToCreate, err := model.NewSegment(model.NewSegmentInput{
		Slug:        segmentSlug,
		WorkspaceID: workspaceID,
		Conditions:  segmentConditions,
	})
	if err != nil {
		return err
	}

	err = uc.segmentRepo.Save(ctx, *segmentToCreate)
	if err != nil {
		return err
	}

	return nil
}

func (uc *CrudSegmentService) UpdateSegment(ctx context.Context, input input.UpdateSegmentInput) (err error) {
	err = input.Validate()
	if err != nil {
		return err
	}
	segmentID := model.SegmentID(input.SegmentID)
	workspaceID := model.WorkspaceID(input.WorkspaceID)
	segmentSlug, err := vo.NewSlug(input.Slug)
	if err != nil {
		return err
	}
	conditions, err := parseConditions(input.Conditions)
	if err != nil {
		return err
	}

	segment, err := uc.segmentRepo.GetByID(ctx, segmentID, workspaceID)
	if err != nil {
		return err
	}

	segment.SetSlug(segmentSlug)
	segment.SetConditions(conditions)

	err = uc.segmentRepo.Save(ctx, segment)
	if err != nil {
		return err
	}

	return nil
}

func (uc *CrudSegmentService) DeleteSegment(ctx context.Context, input input.DeleteSegmentInput) (err error) {
	err = input.Validate()
	if err != nil {
		return err
	}
	segmentID := model.SegmentID(input.SegmentID)
	workspaceID := model.WorkspaceID(input.WorkspaceID)

	activeCampaigns, err := uc.campaignRepo.FindAll(ctx, workspaceID)
	if err != nil {
		return err
	}

	for _, campaign := range activeCampaigns {
		attachedInSomeCampaignFilter := util.Includes(campaign.Filters(), segmentID)
		attachedInSomeCampaignTrigger := util.Includes(campaign.Triggers(), segmentID)
		if attachedInSomeCampaignFilter || attachedInSomeCampaignTrigger {
			return domain.NewInvalidOperationError("delete-segment", "attached in some campaign")
		}

	}

	err = uc.segmentRepo.Delete(ctx, segmentID, workspaceID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *CrudSegmentService) Get(ctx context.Context, input input.GetSegmentInput) (segment model.Segment, err error) {
	err = input.Validate()
	if err != nil {
		return segment, err
	}
	segmentID := model.SegmentID(input.SegmentID)
	workspaceID := model.WorkspaceID(input.WorkspaceID)

	segment, err = uc.segmentRepo.GetByID(ctx, segmentID, workspaceID)
	if err != nil {
		return segment, err
	}

	return segment, nil
}

func (uc *CrudSegmentService) List(ctx context.Context, input input.ListSegmentInput) (segments []model.Segment, err error) {
	err = input.Validate()
	if err != nil {
		return segments, err
	}
	workspaceID := model.WorkspaceID(input.WorkspaceID)
	segments, err = uc.segmentRepo.List(ctx, workspaceID)
	if err != nil {
		return segments, err
	}

	return segments, nil
}

func parseConditions(conditions []input.Condition) (modelConditions []model.Condition, err error) {
	for _, inputCondition := range conditions {
		eventSlug, _ := vo.NewSlug(inputCondition.EventSlug)
		attrKey, _ := vo.NewDotNotation(inputCondition.AttributeKey)
		modelCondition, err := model.NewCondition(model.NewConditionInput{
			ConditionTarget: model.ConditionTarget(inputCondition.ConditionTarget),
			ConditionType:   model.ConditionType(inputCondition.ConditionType),
			EventSlug:       eventSlug,
			AttributeKey:    attrKey,
			AttributeValue:  inputCondition.AttributeValue,
		})
		if err != nil {
			return modelConditions, err
		}
		modelConditions = append(modelConditions, modelCondition)
	}
	return modelConditions, nil
}
