package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/control/input"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

type CrudWorkspaceUsecase struct {
	repo repository.WorkspaceRepository
}

func NewCrudWorkspaceUsecase(repo repository.WorkspaceRepository) *CrudWorkspaceUsecase {
	return &CrudWorkspaceUsecase{
		repo: repo,
	}
}

func (uc *CrudWorkspaceUsecase) CreateWorkspace(ctx context.Context, input input.CreateWorkspaceInput) (err error) {
	err = input.Validate()
	if err != nil {
		return err
	}
	workspaceSlug, err := vo.NewSlug(input.Slug)
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

func (uc *CrudWorkspaceUsecase) UpdateWorkspace(ctx context.Context, input input.UpdateWorkspaceInput) (err error) {
	err = input.Validate()
	if err != nil {
		return err
	}
	workspaceID := model.WorkspaceID(input.WorkspaceID)
	workspaceSlug, err := vo.NewSlug(input.Slug)
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

func (uc *CrudWorkspaceUsecase) DeleteWorkspace(ctx context.Context, input input.DeleteWorkspaceInput) (err error) {
	err = input.Validate()
	if err != nil {
		return err
	}
	workspaceID := model.WorkspaceID(input.WorkspaceID)
	err = uc.repo.Delete(ctx, workspaceID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *CrudWorkspaceUsecase) Get(ctx context.Context, input input.GetWorkspaceInput) (workspace model.Workspace, err error) {
	err = input.Validate()
	if err != nil {
		return workspace, err
	}
	workspaceID := model.WorkspaceID(input.WorkspaceID)

	workspace, err = uc.repo.GetByID(ctx, workspaceID)
	if err != nil {
		return workspace, err
	}

	return workspace, nil
}

func (uc *CrudWorkspaceUsecase) List(ctx context.Context) (workspaces []model.Workspace, err error) {
	workspaces, err = uc.repo.List(ctx)
	if err != nil {
		return workspaces, err
	}

	return workspaces, nil
}
