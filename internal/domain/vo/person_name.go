package vo

import (
	"strings"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type PersonName string

func NewPersonName(name string) (personName PersonName, err domain.DomainError) {
	if util.IsEmpty(name) {
		return personName, domain.NewInvalidEmptyParamError("name")
	}
	return PersonName(name), nil
}

func (vo PersonName) String() string {
	return string(vo)
}

func (vo *PersonName) GetFirstName() string {
	return strings.Title(strings.Split(vo.String(), " ")[0])
}
