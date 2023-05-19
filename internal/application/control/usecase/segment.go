package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/control/command"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

type SegmentUsecase struct {
	repo repository.SegmentRepository
}

func NewSegmentUsecase(repo repository.SegmentRepository) *SegmentUsecase {
	return &SegmentUsecase{
		repo: repo,
	}
}

func (uc *SegmentUsecase) CreateSegment(ctx context.Context, command command.CreateSegmentCommand) (err error) {
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

	workspaceToCreate, err := model.NewSegment(model.NewSegmentInput{
		Slug:        segmentSlug,
		WorkspaceID: workspaceID,
		Conditions:  segmentConditions,
	})
	if err != nil {
		return err
	}

	err = uc.repo.Save(ctx, *workspaceToCreate)
	if err != nil {
		return err
	}

	return nil
}

func (uc *SegmentUsecase) UpdateSegment(ctx context.Context, command command.UpdateSegmentCommand) (err error) {
	err = command.Validate()
	if err != nil {
		return err
	}
	workspaceID := model.SegmentID(command.SegmentID)
	workspaceSlug, err := vo.NewSlug(command.Slug)
	if err != nil {
		return err
	}
	workspace, err := uc.repo.GetByID(ctx, workspaceID)
	if err != nil {
		return err
	}

	workspace.SetSlug(workspaceSlug)

	err = uc.repo.Save(ctx, workspace)
	if err != nil {
		return err
	}

	return nil
}

func (uc *SegmentUsecase) DeleteSegment(ctx context.Context, command command.DeleteSegmentCommand) (err error) {
	err = command.Validate()
	if err != nil {
		return err
	}
	workspaceID := model.SegmentID(command.SegmentID)
	err = uc.repo.Delete(ctx, workspaceID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *SegmentUsecase) Get(ctx context.Context, command command.GetSegmentCommand) (workspace model.Segment, err error) {
	err = command.Validate()
	if err != nil {
		return workspace, err
	}
	workspaceID := model.SegmentID(command.SegmentID)

	workspace, err = uc.repo.GetByID(ctx, workspaceID)
	if err != nil {
		return workspace, err
	}

	return workspace, nil
}

func (uc *SegmentUsecase) List(ctx context.Context) (workspaces []model.Segment, err error) {
	workspaces, err = uc.repo.List(ctx)
	if err != nil {
		return workspaces, err
	}

	return workspaces, nil
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
