package types

import "github.com/itera-io/taikungoclient/models"

var ShowbackTypes = map[string]interface{}{
	"count": models.PrometheusType(100),
	"sum":   models.PrometheusType(200),
}

func GetShowbackType(showbackType string) models.PrometheusType {
	return ShowbackTypes[showbackType].(models.PrometheusType)
}

var ShowbackKinds = map[string]interface{}{
	"general":  models.ShowbackType(100),
	"external": models.ShowbackType(200),
}

func GetShowbackKind(showbackKind string) models.ShowbackType {
	return ShowbackKinds[showbackKind].(models.ShowbackType)
}
