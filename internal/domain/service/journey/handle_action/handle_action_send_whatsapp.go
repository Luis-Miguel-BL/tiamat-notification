package handle_action

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/gateway"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
)

func handleSendWhatsapp(ctx context.Context, gatewayManager gateway.GatewayManager, customer model.Customer, action model.Action) (status model.StepJourneyStatus, nextActionID model.ActionID, err error) {
	return
}
