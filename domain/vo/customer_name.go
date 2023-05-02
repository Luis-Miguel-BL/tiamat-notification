package vo

import (
	"fmt"
	"strings"

	"github.com/Luis-Miguel-BL/tiamat-notification/domain/util"
)

type CustomerName string

func (vo CustomerName) String() string {
	return string(vo)
}

func (vo *CustomerName) Validate() error {
	if util.IsEmpty(vo.String()) {
		return fmt.Errorf("customer_name is empty")
	}
	return nil
}

func (vo *CustomerName) GetFirstName() string {
	return strings.Title(strings.Split(vo.String(), " ")[0])
}
