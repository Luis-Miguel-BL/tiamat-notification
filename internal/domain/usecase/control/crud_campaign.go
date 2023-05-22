package usecase

import (
	"context"
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/usecase/control/command"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type CrudCampaignUsecase struct {
	campaignRepo repository.CampaignRepository
}

func NewCrudCampaignUsecase(campaignRepo repository.CampaignRepository) *CrudCampaignUsecase {
	return &CrudCampaignUsecase{
		campaignRepo: campaignRepo,
	}
}

func (uc *CrudCampaignUsecase) CreateCampaign(ctx context.Context, command command.CreateCampaignCommand) (err error) {
	err = command.Validate()
	if err != nil {
		return err
	}

	campaignSlug, err := vo.NewSlug(command.Slug)
	if err != nil {
		return err
	}
	workspaceID := model.WorkspaceID(command.WorkspaceID)
	retriggerDelay := time.Second * time.Duration(command.RetriggerDelayInSeconds)

	actions, err := parseActions(command.Actions)
	if err != nil {
		return err
	}

	campaignToCreate, err := model.NewCampaign(model.NewCampaignInput{
		WorkspaceID:    workspaceID,
		Slug:           campaignSlug,
		RetriggerDelay: retriggerDelay,

		Conditions: campaignConditions,
	})
	if err != nil {
		return err
	}

	err = uc.campaignRepo.Save(ctx, *campaignToCreate)
	if err != nil {
		return err
	}

	return nil
}

func (uc *CrudCampaignUsecase) UpdateCampaign(ctx context.Context, command command.UpdateCampaignCommand) (err error) {
	err = command.Validate()
	if err != nil {
		return err
	}
	campaignID := model.CampaignID(command.CampaignID)
	workspaceID := model.WorkspaceID(command.WorkspaceID)
	campaignSlug, err := vo.NewSlug(command.Slug)
	if err != nil {
		return err
	}
	conditions, err := parseConditions(command.Conditions)
	if err != nil {
		return err
	}

	campaign, err := uc.campaignRepo.GetByID(ctx, campaignID, workspaceID)
	if err != nil {
		return err
	}

	campaign.SetSlug(campaignSlug)
	campaign.SetConditions(conditions)

	err = uc.campaignRepo.Save(ctx, campaign)
	if err != nil {
		return err
	}

	return nil
}

func (uc *CrudCampaignUsecase) DeleteCampaign(ctx context.Context, command command.DeleteCampaignCommand) (err error) {
	err = command.Validate()
	if err != nil {
		return err
	}
	campaignID := model.CampaignID(command.CampaignID)
	workspaceID := model.WorkspaceID(command.WorkspaceID)

	activeCampaigns, err := uc.campaignRepo.FindActiveCampaigns(ctx, workspaceID)
	if err != nil {
		return err
	}

	for _, campaign := range activeCampaigns {
		attachedInSomeCampaignFilter := util.Includes[model.CampaignID](campaign.Filters(), campaignID)
		attachedInSomeCampaignTrigger := util.Includes[model.CampaignID](campaign.Triggers(), campaignID)
		if attachedInSomeCampaignFilter || attachedInSomeCampaignTrigger {
			return domain.NewInvalidOperationError("delete-campaign", "attached in some campaign")
		}

	}

	err = uc.campaignRepo.Delete(ctx, campaignID, workspaceID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *CrudCampaignUsecase) Get(ctx context.Context, command command.GetCampaignCommand) (campaign model.Campaign, err error) {
	err = command.Validate()
	if err != nil {
		return campaign, err
	}
	campaignID := model.CampaignID(command.CampaignID)
	workspaceID := model.WorkspaceID(command.WorkspaceID)

	campaign, err = uc.campaignRepo.GetByID(ctx, campaignID, workspaceID)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (uc *CrudCampaignUsecase) List(ctx context.Context, command command.ListCampaignCommand) (campaigns []model.Campaign, err error) {
	err = command.Validate()
	if err != nil {
		return campaigns, err
	}
	workspaceID := model.WorkspaceID(command.WorkspaceID)
	campaigns, err = uc.campaignRepo.List(ctx, workspaceID)
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func parseActions(actions []command.Action) (modelActions []model.Action, err error) {
	for _, commandAction := range actions {
		actionID := model.ActionID(commandAction.ActionID)
		nextActionID := model.ActionID(commandAction.NextActionID)
		slug, err := vo.NewSlug(commandAction.Slug)
		if err != nil {
			return modelActions, err
		}

		attrKey, _ := vo.NewDotNotation(commandAction.AttributeKey)
		modelAction, err := model.NewAction(model.NewActionInput{
			ActionID:     actionID,
			Slug:         slug,
			ActionType:   model.ActionType(commandAction.ActionType),
			NextActionID: nextActionID,
		})
		if err != nil {
			return modelActions, err
		}
		modelActions = append(modelActions, modelAction)
	}
	return modelActions, nil
}
