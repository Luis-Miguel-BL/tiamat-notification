package match_condition

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
)

func isMatchByHasBeenPerformedCondition(condition model.Condition, customer model.SerializedCustomer) (match bool) {
	_, match = customer.Events[condition.EventSlug]
	return match
}
