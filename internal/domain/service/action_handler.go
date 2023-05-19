package service

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/gateway"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/service/action_handler"
)

type ActionHandlerService interface {
	HandleAction(ctx context.Context, customer *model.Customer, action model.Action, campaignFilters []model.Segment) (err error)
}

type actionHandlerService struct {
	gatewayManager gateway.GatewayManager
}

func NewActionHandlerService(gatewayManager gateway.GatewayManager) ActionHandlerService {
	return &actionHandlerService{
		gatewayManager: gatewayManager,
	}
}

type Handle func(context.Context, gateway.GatewayManager, *model.Customer, model.Action) (model.ActionTriggeredStatus, model.ActionID, error)

var mapHandler = map[model.BehaviorType]Handle{
	model.BehaviorTypeSendEmail:    action_handler.HandleSendEmail,
	model.BehaviorTypeSendSMS:      action_handler.HandleSendSMS,
	model.BehaviorTypeSendWhatsapp: action_handler.HandleSendWhatsapp,
	model.BehaviorTypeWaitFor:      action_handler.HandleWaitFor,
	model.BehaviorTypeIfAttribute:  action_handler.HandleIfAttribute,
	model.BehaviorTypeRandom:       action_handler.HandleRandom,
	model.BehaviorTypeSplit:        action_handler.HandleSplit,
}

func (s *actionHandlerService) HandleAction(ctx context.Context, customer *model.Customer, action model.Action, campaignFilters []model.Segment) (err error) {
	matchFilters, err := matchFilters(campaignFilters, customer)
	if err != nil {
		return err
	}
	if matchFilters {
		err := customer.FinishActionTriggered(action, model.ActionTriggeredStatusFilterMatch, "")
		if err != nil {
			return err
		}
	}
	status, nextActionID, err := mapHandler[action.BehaviorType()](ctx, s.gatewayManager, customer, action)

	err = customer.FinishActionTriggered(action, status, nextActionID)
	if err != nil {
		return err
	}

	return nil
}

func matchFilters(campaignFilters []model.Segment, customer *model.Customer) (isMatchAll bool, err error) {
	isMatchAll = true
	for _, segment := range campaignFilters {
		for _, condition := range segment.Conditions() {
			if !condition.IsMatch(customer.Serialize()) {
				return false, nil
			}
		}

		satisfiedSegment, err := model.NewSatisfiedSegment(
			model.NewSatisfiedSegmentInput{
				CustomerID:  customer.CustomerID(),
				WorkspaceID: customer.WorkspaceID(),
				SegmentID:   segment.SegmentID(),
			},
		)
		if err != nil {
			return false, err
		}
		customer.AppendSatisfiedSegment(*satisfiedSegment)
	}
	return isMatchAll, nil
}
