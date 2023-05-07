package model

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

var AggregateTypeWorkspace = domain.AggregateType("workspace")

type WorkspaceID string

func NewWorkspaceID(workspaceID string) WorkspaceID {
	return WorkspaceID(workspaceID)
}

type Workspace struct {
	*domain.AggregateRoot
	WorkspaceID WorkspaceID
	Slug        vo.Slug
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewWorkspace(workspaceID string, slug string) (workspace *Workspace, err domain.DomainError) {
	if workspaceID == "" {
		return workspace, domain.NewInvalidEmptyParamError("WorkspaceID")
	}
	return &Workspace{
		WorkspaceID: WorkspaceID(workspaceID),
		Slug:        vo.Slug(slug),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
