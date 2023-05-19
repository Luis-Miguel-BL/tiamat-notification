package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/control/command"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

type WorkspaceUsecase struct {
	repo repository.WorkspaceRepository
}

func NewWorkspaceUsecase(repo repository.WorkspaceRepository) *WorkspaceUsecase {
	return &WorkspaceUsecase{
		repo: repo,
	}
}

func (uc *WorkspaceUsecase) CreateWorkspace(ctx context.Context, command command.CreateWorkspaceCommand) (err error) {
	err = command.Validate()
	if err != nil {
		return err
	}
	workspaceSlug, err := vo.NewSlug(command.Slug)
	if err != nil {
		return err
	}
	workspaceToCreate, err := model.NewWorkspace(workspaceSlug)
	if err != nil {
		return err
	}

	err = uc.repo.Save(ctx, *workspaceToCreate)
	if err != nil {
		return err
	}

	return nil
}

func (uc *WorkspaceUsecase) UpdateWorkspace(ctx context.Context, command command.UpdateWorkspaceCommand) (err error) {
	err = command.Validate()
	if err != nil {
		return err
	}
	workspaceID := model.WorkspaceID(command.WorkspaceID)
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

func (uc *WorkspaceUsecase) DeleteWorkspace(ctx context.Context, command command.DeleteWorkspaceCommand) (err error) {
	err = command.Validate()
	if err != nil {
		return err
	}
	workspaceID := model.WorkspaceID(command.WorkspaceID)
	err = uc.repo.Delete(ctx, workspaceID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *WorkspaceUsecase) Get(ctx context.Context, command command.GetWorkspaceCommand) (workspace model.Workspace, err error) {
	err = command.Validate()
	if err != nil {
		return workspace, err
	}
	workspaceID := model.WorkspaceID(command.WorkspaceID)

	workspace, err = uc.repo.GetByID(ctx, workspaceID)
	if err != nil {
		return workspace, err
	}

	return workspace, nil
}

func (uc *WorkspaceUsecase) List(ctx context.Context) (workspaces []model.Workspace, err error) {
	workspaces, err = uc.repo.List(ctx)
	if err != nil {
		return workspaces, err
	}

	return workspaces, nil
}
