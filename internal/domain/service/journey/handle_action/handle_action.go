package handle_action

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/gateway"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
)

func Handle(ctx context.Context, gatewayManager gateway.GatewayManager, action model.Action, customer model.Customer) (status model.StepJourneyStatus, nextActionID model.ActionID, err error) {

	type handle func(context.Context, gateway.GatewayManager, model.Customer, model.Action) (model.StepJourneyStatus, model.ActionID, error)

	var mapHandler = map[model.BehaviorType]handle{
		model.BehaviorTypeSendEmail:    handleSendEmail,
		model.BehaviorTypeSendSMS:      handleSendSMS,
		model.BehaviorTypeSendWhatsapp: handleSendWhatsapp,
		model.BehaviorTypeWaitFor:      handleWaitFor,
		model.BehaviorTypeWaitUntil:    handleWaitUntil,
		model.BehaviorTypeIfAttribute:  handleIfAttribute,
		model.BehaviorTypeRandom:       handleRandom,
		model.BehaviorTypeSplit:        handleSplit,
	}

	return mapHandler[action.BehaviorType()](ctx, gatewayManager, customer, action)
}
