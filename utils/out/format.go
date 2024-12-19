package out

import (
	"fmt"
	"math"
	"strings"

	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/types"
)

const trimmedValueSuffix = "..."

func trimCellValue(value interface{}) interface{} {
	if !config.ShowLargeValues {
		if str, isString := value.(string); isString {
			if len(str) > config.MaxCellWidth {
				str = str[:(config.MaxCellWidth - len(trimmedValueSuffix))]
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
			return field.NotAvailable
		}

		dateTime = strings.Replace(dateTime, "T", " ", 1)
		dateTime = strings.Replace(dateTime, "Z", "", 1)

		return dateTime
	}

	return field.NotAvailable
}

// Display true/false as Locked/Unlocked
func FormatLockStatus(v interface{}) string {
	if lockStatus, ok := v.(bool); ok {
		if lockStatus {
			return "Locked"
		}

		return "Unlocked"
	}

	return field.NotAvailable
}

// If not available, display N/A
func FormatProjectHealth(v interface{}) string {
	if health, ok := v.(string); ok {
		if health == "None" {
			return field.NotAvailable
		}

		return health
	}

	return field.NotAvailable
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
		case "google":
			return "Google"
		case "vsphere":
			return "vSphere"
		case "proxmox":
			return "Proxmox"
		case "zadara":
			return "Zadara"
		case "zededa":
			return "Zededa"
		case "openshift":
			return "Openshift"
		case "tanzu":
			return "Tanzu"
		default:
			return strings.ToLower(cloudType)
		}
	}

	return field.NotAvailable
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

	return field.NotAvailable
}

// Format Bytes as GiB
func FormatBToGiB(v interface{}) string {
	if bytes, ok := v.(float64); ok {
		var jsMaxSafeInteger float64 = 9007199254740991
		if bytes == jsMaxSafeInteger {
			return field.NotAvailable
		}

		return fmt.Sprintf("%d GiB", int(bytes/math.Pow(1024, 3)))
	}

	return field.NotAvailable
}

// Format number
func FormatNumber(v interface{}) string {
	if number, ok := v.(float64); ok {
		var jsMaxSafeInteger float64 = 9007199254740991
		if number == jsMaxSafeInteger {
			return field.NotAvailable
		}

		return fmt.Sprint(number)
	}

	return field.NotAvailable
}

// Format number as integer
func FormatNumberInteger(v interface{}) string {
	if number, ok := v.(float64); ok {
		var jsMaxSafeInteger float64 = 9007199254740991
		if number == jsMaxSafeInteger {
			return field.NotAvailable
		}

		return fmt.Sprint(int64(number))
	}

	return field.NotAvailable
}

func FormatNumberAddGibString(v interface{}) string {
	if number, ok := v.(float64); ok {
		var jsMaxSafeInteger float64 = 9007199254740991
		if number == jsMaxSafeInteger {
			return field.NotAvailable
		}

		return fmt.Sprintf("%d Gib", int64(number))
	}

	return field.NotAvailable
}

func FormatAvailabilityZones(v interface{}) string {
	switch v.(type) {
	case int:
		return fmt.Sprint(v)
	case string:
		return fmt.Sprint(v)
	default:
		return field.NotAvailable
	}
}

// Format resource ID
func FormatID(v interface{}) string {
	if id, ok := v.(string); ok && id != "0" {
		return id
	}

	return field.NotAvailable
}

// Format RAM from Bytes to GibiBytes
func FormatRAM(v interface{}) string {
	if size, ok := v.(float64); ok && size != 0 {
		GibiBytes := v.(float64) / float64(1073741824)
		if GibiBytes > 1 && GibiBytes < 3000 {
			return fmt.Sprintf("%v GiB", GibiBytes)
		}
	}
	return field.NotAvailable
}

// Format Slack channel
func FormatSlackChannel(v interface{}) string {
	if channel, ok := v.(string); ok {
		return fmt.Sprintf("#%s", channel)
	}

	return field.NotAvailable
}

// Format string as all caps
func FormatStringUpper(v interface{}) string {
	if str, ok := v.(string); ok {
		return strings.ToUpper(str)
	}

	return field.NotAvailable
}

// Format standalone VM tag list
func FormatVMTags(v interface{}) (str string) {
	str = field.NotAvailable

	if tags, ok := v.([]interface{}); ok {
		var stringBuilder strings.Builder

		if tagCount := len(tags); tagCount != 0 {
			tag, tagFormatIsValid := formatVMTag(tags[0])
			if tagFormatIsValid {
				stringBuilder.WriteString(tag)
			}

			for i := 1; i < tagCount && tagFormatIsValid; i++ {
				tag, tagFormatIsValid = formatVMTag(tags[i])

				stringBuilder.WriteString(",")
				stringBuilder.WriteString(tag)
			}

			if tagFormatIsValid {
				str = stringBuilder.String()
			}
		}
	}

	return
}

func formatVMTag(v interface{}) (str string, ok bool) {
	var stringBuilder strings.Builder

	tagMap, ok := v.(map[string]interface{})
	if !ok {
		return
	}

	key, ok := getTagValue(tagMap, "key")
	if !ok {
		return
	}

	stringBuilder.WriteString(key)
	stringBuilder.WriteString("=")

	value, ok := getTagValue(tagMap, "value")
	if !ok {
		return
	}

	stringBuilder.WriteString(value)

	str = stringBuilder.String()

	return
}

func getTagValue(tagMap map[string]interface{}, key string) (valueString string, ok bool) {
	value, ok := tagMap[key]
	if ok {
		valueString, ok = value.(string)
	}

	return
}

// FormatRepoName extracts the "name" field from the "repository" map.
func FormatRepoName(repository interface{}) string {
	// Assert the repository to a map[string]interface{} to handle dynamic keys.
	if repoMap, ok := repository.(map[string]interface{}); ok {
		if name, exists := repoMap["name"]; exists {
			if nameStr, ok := name.(string); ok {
				return nameStr
			}
		}
	}
	return ""
}

func FormatDisabled(isBound interface{}) string {
	if disabledBool, ok := isBound.(bool); ok {
		if disabledBool {
			return "Enabled"
		} else {
			return "Disabled"
		}
	}
	return ""
}
