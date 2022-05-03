package utils

import (
	"strings"

	"github.com/google/uuid"
)

func GetUUID() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}
