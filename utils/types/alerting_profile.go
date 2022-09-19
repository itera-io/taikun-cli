package types

import (
	"github.com/itera-io/taikun-cli/utils/gmap"
	"github.com/itera-io/taikungoclient/models"
)

const (
	AlertingIntegrationTypeTeams = "microsoftteams"
)

var AlertingIntegrationTypes = gmap.New(
	map[string]interface{}{
		"opsgenie":                   models.AlertingIntegrationType(100),
		"pagerduty":                  models.AlertingIntegrationType(200),
		"splunk":                     models.AlertingIntegrationType(300),
		AlertingIntegrationTypeTeams: models.AlertingIntegrationType(400),
	},
)

func GetAlertingIntegrationType(integrationType string) models.AlertingIntegrationType {
	model, _ := AlertingIntegrationTypes.Get(integrationType).(models.AlertingIntegrationType)
	return model
}

var AlertingReminders = gmap.New(
	map[string]interface{}{
		"halfhour": models.AlertingReminder(100),
		"hourly":   models.AlertingReminder(200),
		"daily":    models.AlertingReminder(300),
		"none":     models.AlertingReminder(-1),
	},
)

func GetAlertingReminder(reminder string) models.AlertingReminder {
	model, _ := AlertingReminders.Get(reminder).(models.AlertingReminder)
	return model
}
