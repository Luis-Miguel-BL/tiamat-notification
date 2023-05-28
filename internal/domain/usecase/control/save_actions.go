package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/usecase/control/command"
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

func (uc *SaveActionsUsecase) SaveActions(ctx context.Context, command command.SaveActionsCommand) (err error) {
	err = command.Validate()
	if err != nil {
		return err
	}
	workspaceID := model.WorkspaceID(command.WorkspaceID)
	campaignID := model.CampaignID(command.CampaignID)
	firstActionID := model.ActionID(command.FirstActionID)
	actions, err := parseActions(command.Actions)
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

func parseActions(actions []command.Action) (modelActions map[model.ActionID]model.Action, err error) {
	for _, commandAction := range actions {
		actionID := model.ActionID(commandAction.ActionID)
		actionType := model.ActionType(commandAction.ActionType)
		behaviorType := model.BehaviorType(commandAction.BehaviorType)
		behavior := commandAction.Behavior
		slug, err := vo.NewSlug(commandAction.Slug)
		if err != nil {
			return modelActions, err
		}
		nextActionsID := []model.ActionID{}
		for _, nextActionID := range commandAction.NextActionsID {
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
			ActionType:    model.ActionType(commandAction.ActionType),
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
