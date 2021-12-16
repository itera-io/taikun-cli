package types

import "github.com/itera-io/taikungoclient/models"

var AlertingIntegrationTypes = map[string]interface{}{
	"opsgenie":       100,
	"pagerduty":      200,
	"splunk":         300,
	"microsoftteams": 400,
}

func GetAlertingIntegrationType(integrationType string) models.AlertingIntegrationType {
	return AlertingIntegrationTypes[integrationType].(models.AlertingIntegrationType)
}

var AlertingReminders = map[string]interface{}{
	"halfhour": 100,
	"hourly":   200,
	"daily":    300,
	"none":     -1,
}

func GetAlertingReminder(reminder string) models.AlertingReminder {
	return AlertingReminders[reminder].(models.AlertingReminder)
}
