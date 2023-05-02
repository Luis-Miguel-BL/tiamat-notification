package util

import "strings"

func IsEmpty(value string) bool {
	trimValue := strings.Trim(value, " ")
	return trimValue == ""
}
