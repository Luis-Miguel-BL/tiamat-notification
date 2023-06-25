package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/control/input"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

type SaveActionsUsecase struct {
	campaignRepo repository.CampaignRepository
}

func NewSaveActionsUsecase(campaignRepo repository.CampaignRepository) *SaveActionsUsecase {
	return &SaveActionsUsecase{
		campaignRepo: campaignRepo,
	}
}

func (uc *SaveActionsUsecase) SaveActions(ctx context.Context, input input.SaveActionsInput) (err error) {
	err = input.Validate()
	if err != nil {
		return err
	}
	workspaceID := model.WorkspaceID(input.WorkspaceID)
	campaignID := model.CampaignID(input.CampaignID)
	firstActionID := model.ActionID(input.FirstActionID)
	actions, err := parseActions(input.Actions)
	if err != nil {
		return err
	}

	campaign, err := uc.campaignRepo.GetByID(ctx, campaignID, workspaceID)
	if err != nil {
		return err
	}

	campaign.SetActions(firstActionID, actions)

	err = uc.campaignRepo.Save(ctx, campaign)
	if err != nil {
		return err
	}

	return nil
}

func parseActions(actions []input.Action) (modelActions map[model.ActionID]model.Action, err error) {
	modelActions = make(map[model.ActionID]model.Action)
	for _, inputAction := range actions {
		actionID := model.ActionID(inputAction.ActionID)
		actionType := model.ActionType(inputAction.ActionType)
		behaviorType := model.BehaviorType(inputAction.BehaviorType)
		behavior := inputAction.Behavior
		slug, err := vo.NewSlug(inputAction.Slug)
		if err != nil {
			return modelActions, err
		}
		nextActionsID := []model.ActionID{}
		for _, nextActionID := range inputAction.NextActionsID {
			nextActionsID = append(nextActionsID, model.ActionID(nextActionID))
		}

		_, availableActionType := model.AvailableActionType[actionType]
		if !availableActionType {
			return modelActions, domain.NewInvalidParamError("action-type")
		}
		validateBehavior, availableBehaviorType := model.AvailableBehaviorType[behaviorType]
		if !availableBehaviorType {
			return modelActions, domain.NewInvalidParamError("behavior-type")
		}
		if !validateBehavior(behavior) {
			return modelActions, domain.NewInvalidParamError("behavior")
		}

		modelAction, err := model.NewAction(model.NewActionInput{
			ActionID:      actionID,
			Slug:          slug,
			ActionType:    model.ActionType(inputAction.ActionType),
			NextActionsID: nextActionsID,
			BehaviorType:  behaviorType,
			Behavior:      behavior,
		})
		if err != nil {
			return modelActions, err
		}
		modelActions[modelAction.ActionID()] = modelAction
	}

	for _, modelAction := range modelActions {
		for _, nextActionID := range modelAction.NextActionsID() {
			_, findNextAction := modelActions[nextActionID]
			if !findNextAction {
				return modelActions, domain.NewInvalidParamError("next-action-id")
			}
		}
	}

	return modelActions, nil
}
