package types

import (
	"strings"

	"github.com/itera-io/taikungoclient/models"
)

var ShowbackKinds = map[string]interface{}{
	"general":  models.ShowbackType(100),
	"external": models.ShowbackType(200),
}

func GetShowbackKind(showbackKind string) models.ShowbackType {
	return ShowbackKinds[strings.ToLower(showbackKind)].(models.ShowbackType)
}
