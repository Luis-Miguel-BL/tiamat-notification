package repository

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
)

type SegmentRepository interface {
	Save(ctx context.Context, segment model.Segment) (err error)
	GetByID(ctx context.Context, segmentID model.SegmentID, workspaceID model.WorkspaceID) (segment model.Segment, err error)
	List(ctx context.Context, workspaceID model.WorkspaceID) (segments []model.Segment, err error)
	Delete(ctx context.Context, segmentID model.SegmentID, workspaceID model.WorkspaceID) (err error)
}
