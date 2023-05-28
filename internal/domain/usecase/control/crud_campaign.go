package usecase

import (
	"context"
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/usecase/control/command"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

type CrudCampaignUsecase struct {
	campaignRepo repository.CampaignRepository
	segmentRepo  repository.SegmentRepository
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
	triggers, err := uc.parseSegmentIDs(ctx, command.Triggers, workspaceID)
	if err != nil {
		return err
	}
	filters, err := uc.parseSegmentIDs(ctx, command.Filters, workspaceID)
	if err != nil {
		return err
	}

	campaignToCreate, err := model.NewCampaign(model.NewCampaignInput{
		WorkspaceID:    workspaceID,
		Slug:           campaignSlug,
		RetriggerDelay: retriggerDelay,
		Triggers:       triggers,
		Filters:        filters,
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
	campaignSlug, err := vo.NewSlug(command.Slug)
	if err != nil {
		return err
	}
	campaignID := model.CampaignID(command.CampaignID)
	workspaceID := model.WorkspaceID(command.WorkspaceID)
	retriggerDelay := time.Second * time.Duration(command.RetriggerDelayInSeconds)
	triggers, err := uc.parseSegmentIDs(ctx, command.Triggers, workspaceID)
	if err != nil {
		return err
	}
	filters, err := uc.parseSegmentIDs(ctx, command.Filters, workspaceID)
	if err != nil {
		return err
	}
	campaign, err := uc.campaignRepo.GetByID(ctx, campaignID, workspaceID)
	if err != nil {
		return err
	}

	campaign.SetSlug(campaignSlug)
	campaign.SetRetriggerDelay(retriggerDelay)
	campaign.SetTriggers(triggers)
	campaign.SetFilters(filters)

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
	campaigns, err = uc.campaignRepo.FindAll(ctx, workspaceID)
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (uc *CrudCampaignUsecase) parseSegmentIDs(ctx context.Context, ids []string, workspaceID model.WorkspaceID) (segmentIDs []model.SegmentID, err error) {
	for _, id := range ids {
		segmentID := model.SegmentID(id)
		_, err := uc.segmentRepo.GetByID(ctx, segmentID, workspaceID)
		if err != nil {
			return segmentIDs, err
		}
		segmentIDs = append(segmentIDs, segmentID)
	}
	return segmentIDs, nil
}
