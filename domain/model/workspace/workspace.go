package workspace

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/domain/util"
	"github.com/Luis-Miguel-BL/tiamat-notification/domain/vo"
)

var AggregateType = domain.AggregateType("workspace")

type WorkspaceID string
type Workspace struct {
	*domain.Aggregate
	WorkspaceID WorkspaceID
	Slug        vo.Slug
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewWorkspace(workspaceID string, slug string) (customer *Workspace) {
	return &Workspace{
		WorkspaceID: WorkspaceID(workspaceID),
		Slug:        vo.Slug(slug),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (e *Workspace) Validate() error {
	if util.IsEmpty(string(e.WorkspaceID)) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	if err := e.Slug.Validate(); err != nil {
		return err
	}
	return nil
}
