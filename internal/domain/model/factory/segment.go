package factory

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

type SegmentFactory struct{}

type CreateSegmentInput struct {
	WorkspaceID string
	Slug        string
	Conditions  []model.Condition
}

func (f *SegmentFactory) CreateSegment(input CreateSegmentInput) (segment *model.Segment) {
	if input.Slug == "" {
		input.Slug = "fake-slug"
	}
	if input.WorkspaceID == "" {
		input.WorkspaceID = "fake-workspace-id"
	}

	segment, _ = model.NewSegment(model.NewSegmentInput{
		Slug:        vo.Slug(input.Slug),
		WorkspaceID: model.WorkspaceID(input.WorkspaceID),
		Conditions:  input.Conditions,
	})

	return segment
}
