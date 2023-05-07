package repository

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
)

type SegmentRepository interface {
	Save(ctx context.Context, segment model.Segment) (err error)
	GetByID(ctx context.Context, segmentID model.SegmentID, workspaceID model.WorkspaceID) (segment model.Segment, err error)
}
