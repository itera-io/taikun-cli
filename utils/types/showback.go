package types

import (
	"github.com/itera-io/taikun-cli/utils/gmap"
	"github.com/itera-io/taikungoclient/models"
)

var ShowbackKinds = gmap.New(
	map[string]interface{}{
		"general":  models.ShowbackType(100),
		"external": models.ShowbackType(200),
	},
)

func GetShowbackKind(showbackKind string) models.ShowbackType {
	model, _ := ShowbackKinds.Get(showbackKind).(models.ShowbackType)
	return model
}
