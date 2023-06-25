package repository

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
)

type JourneyRepository interface {
	Save(ctx context.Context, journey model.Journey) (err error)
	GetByID(ctx context.Context, journeyID model.JourneyID, workspaceID model.WorkspaceID) (journey model.Journey, err error)
	FindByCustomerID(ctx context.Context, workspaceID model.WorkspaceID, customerID model.CustomerID) (journeys []model.Journey, err error)
	GetStepJourney(ctx context.Context, workspaceID model.WorkspaceID, stepJourneyID model.StepJourneyID) (stepJourney model.StepJourney, err error)
	SaveStepJourney(ctx context.Context, stepJourney model.StepJourney) (err error)
}
