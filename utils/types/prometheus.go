package types

import (
	"github.com/itera-io/taikun-cli/utils/gmap"
	"github.com/itera-io/taikungoclient/models"
)

var PrometheusTypes = gmap.New(
	map[string]interface{}{
		"count": models.PrometheusType(100),
		"sum":   models.PrometheusType(200),
	},
)

func GetPrometheusType(showbackType string) models.PrometheusType {
	return PrometheusTypes.Get(showbackType).(models.PrometheusType)
}
