package vo

import (
	"fmt"
	"regexp"

	"github.com/Luis-Miguel-BL/tiamat-notification/domain/util"
)

type EmailAddress string

var regexEmail = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

func (vo EmailAddress) String() string {
	return string(vo)
}
func (vo *EmailAddress) Validate() error {
	if util.IsEmpty(vo.String()) {
		return fmt.Errorf("email_address is empty")
	}

	regex, err := regexp.Compile(regexEmail)
	if err != nil {
		return fmt.Errorf("invalid email")
	}

	if !regex.MatchString(vo.String()) {
		return fmt.Errorf("invalid email")
	}
	return nil
}
