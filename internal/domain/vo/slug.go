package vo

import (
	"regexp"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

type Slug string

func NewSlug(slug string) (newSlug Slug, err domain.DomainError) {
	pattern := "^[a-z0-9]+(-[a-z0-9]+)*$"
	re := regexp.MustCompile(pattern)

	if !re.MatchString(slug) {
		return newSlug, domain.NewInvalidParamError("slug")
	}
	return Slug(slug), nil
}

func (vo Slug) String() string {
	return string(vo)
}
