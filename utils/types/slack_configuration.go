package types

import (
	"strings"

	"github.com/itera-io/taikun-cli/utils/gmap"
	"github.com/itera-io/taikungoclient/models"
)

var SlackTypes = gmap.New(
	map[string]interface{}{
		"alert":   models.SlackType(100),
		"general": models.SlackType(200),
	},
)

func GetSlackType(slackType string) models.SlackType {
	return SlackTypes.Get(strings.ToLower(slackType)).(models.SlackType)
}
