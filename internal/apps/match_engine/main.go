package match_engine

import (
	"context"

	usecase "github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/usecase/match"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/usecase/match/input"
)

func Handle(ctx context.Context, matchCustomer usecase.MatchCustomerUsecase, customerID string, workspaceID string) (err error) {

	err = matchCustomer.MatchCustomer(ctx, input.MatchCustomerInput{
		CustomerID:  customerID,
		WorkspaceID: workspaceID,
	})
	if err != nil {
		return err
	}

	return nil
}
