package out

import (
	"fmt"
	"math"
	"strings"

	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/types"
)

const maxColumnWidth = 50
const trimmedValueSuffix = "..."

func trimCellValue(value interface{}) interface{} {
	if !config.ShowLargeValues {
		if str, isString := value.(string); isString {
			if len(str) > maxColumnWidth {
				str = str[:(maxColumnWidth - len(trimmedValueSuffix))]
				str += trimmedValueSuffix
			}
			return str
		}
	}
	return value
}

func resourceIDToString(id interface{}) string {
	if str, isString := id.(string); isString {
		return strings.ReplaceAll(str, "\"", "")
	}
	return fmt.Sprint(id)
}

// Format a datetime string with format '<YYYY>-<MM>-<DD>T<HH>:<MM>:<SS>Z'
func FormatDateTimeString(v interface{}) string {
	if dateTime, ok := v.(string); ok {
		if dateTime == "" {
			return "N/A"
		}
		dateTime = strings.Replace(dateTime, "T", " ", 1)
		dateTime = strings.Replace(dateTime, "Z", "", 1)
		return dateTime
	}
	return "N/A"

}

// Display true/false as Locked/Unlocked
func FormatLockStatus(v interface{}) string {
	if lockStatus, ok := v.(bool); ok {
		if lockStatus {
			return "Locked"
		}
		return "Unlocked"
	}
	return "N/A"
}

// If not availabale, display N/A
func FormatProjectHealth(v interface{}) string {
	if health, ok := v.(string); ok {
		if health == "None" {
			return "N/A"
		}
		return health
	}
	return "N/A"
}

// Capitalize cloud type
func FormatCloudType(v interface{}) string {
	if cloudType, ok := v.(string); ok {
		switch strings.ToLower(cloudType) {
		case "openstack", "os":
			return "OpenStack"
		case "aws", "amazon":
			return "AWS"
		case "azure":
			return "Azure"
		}
	}
	return "N/A"
}

// Format estimated time of completion
func FormatETC(v interface{}) string {
	if etc, ok := v.(string); ok {
		if etcValue, err := types.Atoi32(etc); err == nil {
			if etcValue == 0 {
				return "Under a minute"
			}
			return fmt.Sprintf("%s minutes", etc)
		}
		return etc
	}
	return "N/A"
}

// Format Bytes as GiB
func FormatBToGiB(v interface{}) string {
	if bytes, ok := v.(float64); ok {
		var jsMaxSafeInteger float64 = 9007199254740991
		if bytes == jsMaxSafeInteger {
			return "N/A"
		}
		return fmt.Sprintf("%d GiB", int(bytes/math.Pow(1024, 3)))
	}
	return "N/A"
}

// Format number
func FormatNumber(v interface{}) string {
	if n, ok := v.(float64); ok {
		var jsMaxSafeInteger float64 = 9007199254740991
		if n == jsMaxSafeInteger {
			return "N/A"
		}
		return fmt.Sprint(n)
	}
	return "N/A"
}

// Format resource ID
func FormatID(v interface{}) string {
	if id, ok := v.(string); ok && id != "0" {
		return id
	}
	return "N/A"
}
