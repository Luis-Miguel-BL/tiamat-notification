package util

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

func IsEmpty(value string) bool {
	trimValue := strings.Trim(value, " ")
	return trimValue == ""
}

func NewUUID() (id string) {
	id = uuid.New().String()
	return id
}

func Includes[T any](values []T, targetValue T) bool {
	for _, value := range values {
		if any(value) == any(targetValue) {
			return true
		}
	}
	return false
}
func NewUnixTime(unixDate uint32) time.Time {
	return time.Unix(int64(unixDate), 0)
}
