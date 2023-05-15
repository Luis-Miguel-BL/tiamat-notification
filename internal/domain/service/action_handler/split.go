package action_handler

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/gateway"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
)

func HandleSplit(ctx context.Context, gatewayManager gateway.GatewayManager, customer *model.Customer, action model.Action) (status model.ActionTriggeredStatus, nextActionID model.ActionID, err error) {
	return
}
