package utils

import (
	"github.com/satori/go.uuid"
)

func NewRandomUUIDString() string {
	return uuid.NewV4().String()
}
