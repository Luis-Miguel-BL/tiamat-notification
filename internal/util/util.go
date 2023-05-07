package util

import (
	"strings"

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
