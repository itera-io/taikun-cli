package types

import (
	"github.com/itera-io/taikun-cli/utils/gmap"
)

var OpenstackContinent = gmap.New(
	map[string]interface{}{
		"europe":  "eu",
		"america": "us",
		"asia":    "as",
	},
)

func GetOpenstackContinent(openstackContinent string) interface{} {
	model := OpenstackContinent.Get(openstackContinent)
	return model
}
