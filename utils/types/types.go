package types

import (
	"fmt"
	"strconv"

	"github.com/itera-io/taikungoclient/models"
)

const (
	UnlockedMode = "unlock"
	LockedMode   = "lock"
)

func MapKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func MapContains(m map[string]interface{}, key string) bool {
	_, contains := m[key]
	return contains
}

func UnknownFlagValueError(flag string, received string, expected []string) error {
	return fmt.Errorf("Unknown %s: %s, expected one of %v.", flag, received, expected)
}

var AlertingIntegrationTypes = map[string]interface{}{
	"opsgenie":       models.AlertingIntegrationType(100),
	"pagerduty":      models.AlertingIntegrationType(200),
	"splunk":         models.AlertingIntegrationType(300),
	"microsoftteams": models.AlertingIntegrationType(400),
}

func GetAlertingIntegrationType(integrationType string) models.AlertingIntegrationType {
	return AlertingIntegrationTypes[integrationType].(models.AlertingIntegrationType)
}

var UserRoles = map[string]interface{}{
	"user":    models.UserRole(400),
	"manager": models.UserRole(200),
}

func GetUserRole(role string) models.UserRole {
	return UserRoles[role].(models.UserRole)
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

func Atoi32(str string) (int32, error) {
	res, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(res), nil
}

func GiBToMiB(gibiBytes float64) int32 {
	return int32(gibiBytes * 1024)
}
