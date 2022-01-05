package types

import (
	"strings"

	"github.com/itera-io/taikungoclient/models"
)

var PrometheusTypes = map[string]interface{}{
	"count": models.PrometheusType(100),
	"sum":   models.PrometheusType(200),
}

func GetPrometheusType(showbackType string) models.PrometheusType {
	return PrometheusTypes[strings.ToLower(showbackType)].(models.PrometheusType)
}
