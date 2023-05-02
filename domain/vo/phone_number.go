package vo

import (
	"fmt"

	"github.com/Luis-Miguel-BL/tiamat-notification/domain/util"
)

type PhoneNumber string

func (vo PhoneNumber) String() string {
	return string(vo)
}
func (vo PhoneNumber) Validate() error {
	if util.IsEmpty(vo.String()) {
		return fmt.Errorf("phone_numer is empty")
	}
	//@todo implements regex to validate phone_number
	return nil
}
