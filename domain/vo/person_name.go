package vo

import (
	"fmt"
	"strings"

	"github.com/Luis-Miguel-BL/tiamat-notification/domain/util"
)

type PersonName string

func NewPersonName(name string) PersonName {
	return PersonName(name)
}

func (vo PersonName) String() string {
	return string(vo)
}

func (vo *PersonName) Validate() error {
	if util.IsEmpty(vo.String()) {
		return fmt.Errorf("person_name is empty")
	}
	return nil
}

func (vo *PersonName) GetFirstName() string {
	return strings.Title(strings.Split(vo.String(), " ")[0])
}
