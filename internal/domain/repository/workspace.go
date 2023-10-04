package repository

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
)

type WorkspaceRepository interface {
	Save(ctx context.Context, workspace model.Workspace) (err error)
	GetByID(ctx context.Context, workspaceID model.WorkspaceID) (workspace model.Workspace, err error)
	List(ctx context.Context) (workspaces []model.Workspace, err error)
	Delete(ctx context.Context, workspaceID model.WorkspaceID) (err error)
}
