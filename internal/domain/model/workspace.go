package model

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

var AggregateTypeWorkspace = domain.AggregateType("workspace")

type WorkspaceID string

type Workspace struct {
	*domain.AggregateRoot
	workspaceID WorkspaceID
	slug        vo.Slug
	createdAt   time.Time
	updatedAt   time.Time
}

func NewWorkspace(slug vo.Slug) (workspace *Workspace, err domain.DomainError) {
	workspaceID := WorkspaceID(util.NewUUID())
	return &Workspace{
		AggregateRoot: domain.NewAggregateRoot(AggregateTypeWorkspace, domain.AggregateID(workspaceID)),
		workspaceID:   workspaceID,
		slug:          slug,
		createdAt:     time.Now(),
		updatedAt:     time.Now(),
	}, nil
}

func (e *Workspace) WorkspaceID() WorkspaceID {
	return e.workspaceID
}
func (e *Workspace) Slug() vo.Slug {
	return e.slug
}
func (e *Workspace) SetSlug(slug vo.Slug) {
	e.updatedAt = time.Now()
	e.slug = slug
}
