package vo

import (
	"fmt"
	"regexp"
)

type Slug string

func (vo Slug) String() string {
	return string(vo)
}
func (vo *Slug) Validate() error {
	pattern := "^[a-z0-9]+(-[a-z0-9]+)*$"
	re := regexp.MustCompile(pattern)

	if !re.MatchString(vo.String()) {
		return fmt.Errorf("invalid slug")
	}
	return nil
}
