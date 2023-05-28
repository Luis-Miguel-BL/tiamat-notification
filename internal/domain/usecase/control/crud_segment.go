package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/usecase/control/command"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type CrudSegmentUsecase struct {
	segmentRepo  repository.SegmentRepository
	campaignRepo repository.CampaignRepository
}

func NewCrudSegmentUsecase(segmentRepo repository.SegmentRepository, campaignRepo repository.CampaignRepository) *CrudSegmentUsecase {
	return &CrudSegmentUsecase{
		segmentRepo:  segmentRepo,
		campaignRepo: campaignRepo,
	}
}

func (uc *CrudSegmentUsecase) CreateSegment(ctx context.Context, command command.CreateSegmentCommand) (err error) {
	err = command.Validate()
	if err != nil {
		return err
	}

	segmentSlug, err := vo.NewSlug(command.Slug)
	if err != nil {
		return err
	}
	workspaceID := model.WorkspaceID(command.WorkspaceID)

	segmentConditions, err := parseConditions(command.Conditions)
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

func (uc *CrudSegmentUsecase) UpdateSegment(ctx context.Context, command command.UpdateSegmentCommand) (err error) {
	err = command.Validate()
	if err != nil {
		return err
	}
	segmentID := model.SegmentID(command.SegmentID)
	workspaceID := model.WorkspaceID(command.WorkspaceID)
	segmentSlug, err := vo.NewSlug(command.Slug)
	if err != nil {
		return err
	}
	conditions, err := parseConditions(command.Conditions)
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

func (uc *CrudSegmentUsecase) DeleteSegment(ctx context.Context, command command.DeleteSegmentCommand) (err error) {
	err = command.Validate()
	if err != nil {
		return err
	}
	segmentID := model.SegmentID(command.SegmentID)
	workspaceID := model.WorkspaceID(command.WorkspaceID)

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

func (uc *CrudSegmentUsecase) Get(ctx context.Context, command command.GetSegmentCommand) (segment model.Segment, err error) {
	err = command.Validate()
	if err != nil {
		return segment, err
	}
	segmentID := model.SegmentID(command.SegmentID)
	workspaceID := model.WorkspaceID(command.WorkspaceID)

	segment, err = uc.segmentRepo.GetByID(ctx, segmentID, workspaceID)
	if err != nil {
		return segment, err
	}

	return segment, nil
}

func (uc *CrudSegmentUsecase) List(ctx context.Context, command command.ListSegmentCommand) (segments []model.Segment, err error) {
	err = command.Validate()
	if err != nil {
		return segments, err
	}
	workspaceID := model.WorkspaceID(command.WorkspaceID)
	segments, err = uc.segmentRepo.List(ctx, workspaceID)
	if err != nil {
		return segments, err
	}

	return segments, nil
}

func parseConditions(conditions []command.Condition) (modelConditions []model.Condition, err error) {
	for _, commandCondition := range conditions {
		eventSlug, _ := vo.NewSlug(commandCondition.EventSlug)
		attrKey, _ := vo.NewDotNotation(commandCondition.AttributeKey)
		modelCondition, err := model.NewCondition(model.NewConditionInput{
			ConditionTarget: model.ConditionTarget(commandCondition.ConditionTarget),
			ConditionType:   model.ConditionType(commandCondition.ConditionType),
			EventSlug:       eventSlug,
			AttributeKey:    attrKey,
			AttributeValue:  commandCondition.AttributeValue,
		})
		if err != nil {
			return modelConditions, err
		}
		modelConditions = append(modelConditions, modelCondition)
	}
	return modelConditions, nil
}
