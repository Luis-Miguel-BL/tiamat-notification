package condition

import "github.com/Luis-Miguel-BL/tiamat-notification/domain/model/segment"

const SegmentConditionTypeMoreThan segment.ConditionType = "contains"

type MoreThan struct {
	*segment.ConditionBase
}

func (c *MoreThan) IsMatch(customer *segment.SerializedCustomer) (match bool, err error) {
	return
}
