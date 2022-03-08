package types

import (
	"github.com/itera-io/taikun-cli/utils/gmap"
)

var GoogleImageTypes = gmap.New(
	map[string]interface{}{
		"all":     "all",
		"ubuntu":  "ubuntu",
		"debian":  "debian",
		"windows": "windows",
	},
)

func GetGoogleImageType(googleImageType string) string {
	name, _ := GoogleImageTypes.Get(googleImageType).(string)
	return name
}
