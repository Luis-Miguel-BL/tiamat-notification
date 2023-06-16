package match_condition

import (
	"fmt"
	"strconv"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
)

func isMatchByMoreThanCondition(condition model.Condition, customer model.SerializedCustomer) (match bool) {
	attributeValue, ok := getCustomerAttributeValue(condition, customer)
	if !ok {
		return false
	}
	intValue, err := strconv.ParseFloat(fmt.Sprintf("%v", attributeValue), 64)
	if err != nil {
		return false
	}

	intConditionValue, err := strconv.ParseFloat(fmt.Sprintf("%v", condition.AttributeValue), 64)
	if err != nil {
		return false
	}
	return intValue > intConditionValue
}
