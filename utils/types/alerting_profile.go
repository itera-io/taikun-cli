package types

import "github.com/itera-io/taikungoclient/models"

var AlertingIntegrationTypes = map[string]interface{}{
	"opsgenie":       models.AlertingIntegrationType(100),
	"pagerduty":      models.AlertingIntegrationType(200),
	"splunk":         models.AlertingIntegrationType(300),
	"microsoftteams": models.AlertingIntegrationType(400),
}

func GetAlertingIntegrationType(integrationType string) models.AlertingIntegrationType {
	return AlertingIntegrationTypes[integrationType].(models.AlertingIntegrationType)
}

var AlertingReminders = map[string]interface{}{
	"halfhour": models.AlertingReminder(100),
	"hourly":   models.AlertingReminder(200),
	"daily":    models.AlertingReminder(300),
	"none":     models.AlertingReminder(-1),
}

func GetAlertingReminder(reminder string) models.AlertingReminder {
	return AlertingReminders[reminder].(models.AlertingReminder)
}
