package types

import (
	"github.com/itera-io/taikun-cli/utils/gmap"
	taikuncore "github.com/itera-io/taikungoclient/client"
)

const (
	AlertingIntegrationTypeTeams = "microsoftteams"
)

var AlertingIntegrationTypes = gmap.New(
	map[string]interface{}{
		"opsgenie":                   taikuncore.ALERTINGINTEGRATIONTYPE_OPSGENIE,
		"pagerduty":                  taikuncore.ALERTINGINTEGRATIONTYPE_PAGERDUTY,
		"splunk":                     taikuncore.ALERTINGINTEGRATIONTYPE_SPLUNK,
		AlertingIntegrationTypeTeams: taikuncore.ALERTINGINTEGRATIONTYPE_MICROSOFT_TEAMS,
	},
)

func GetAlertingIntegrationType(integrationType string) *taikuncore.AlertingIntegrationType {
	model, _ := AlertingIntegrationTypes.Get(integrationType).(taikuncore.AlertingIntegrationType)
	return &model
}

var AlertingReminders = gmap.New(
	map[string]interface{}{
		"halfhour": taikuncore.ALERTINGREMINDER_HALF_HOUR,
		"hourly":   taikuncore.ALERTINGREMINDER_HOURLY,
		"daily":    taikuncore.ALERTINGREMINDER_DAILY,
		"none":     taikuncore.ALERTINGREMINDER_NONE,
	},
)

func GetAlertingReminder(reminder string) *taikuncore.AlertingReminder {
	model, _ := AlertingReminders.Get(reminder).(taikuncore.AlertingReminder)
	return &model
}

//import (
//	"github.com/itera-io/taikun-cli/utils/gmap"
//	"github.com/itera-io/taikungoclient/models"
//)
//
//const (
//	AlertingIntegrationTypeTeams = "microsoftteams"
//)
//
//var AlertingIntegrationTypes = gmap.New(
//	map[string]interface{}{
//		"opsgenie":                   models.AlertingIntegrationType(100),
//		"pagerduty":                  models.AlertingIntegrationType(200),
//		"splunk":                     models.AlertingIntegrationType(300),
//		AlertingIntegrationTypeTeams: models.AlertingIntegrationType(400),
//	},
//)
//
//func GetAlertingIntegrationType(integrationType string) models.AlertingIntegrationType {
//	model, _ := AlertingIntegrationTypes.Get(integrationType).(models.AlertingIntegrationType)
//	return model
//}
//
//var AlertingReminders = gmap.New(
//	map[string]interface{}{
//		"halfhour": models.AlertingReminder(100),
//		"hourly":   models.AlertingReminder(200),
//		"daily":    models.AlertingReminder(300),
//		"none":     models.AlertingReminder(-1),
//	},
//)
//
//func GetAlertingReminder(reminder string) models.AlertingReminder {
//	model, _ := AlertingReminders.Get(reminder).(models.AlertingReminder)
//	return model
//}
