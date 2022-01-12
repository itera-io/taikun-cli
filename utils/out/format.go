package out

import (
	"strings"

	"github.com/itera-io/taikun-cli/config"
)

func formatFieldName(fieldName string) string {
	var stringBuilder strings.Builder
	previous := fieldName[0]
	stringBuilder.WriteByte(toUpper(fieldName[0]))
	for i := 1; i < len(fieldName); i++ {
		if isLowerCase(previous) && isUpperCase(fieldName[i]) {
			stringBuilder.WriteByte(' ')
		}
		stringBuilder.WriteByte(fieldName[i])
		previous = fieldName[i]
	}
	return stringBuilder.String()
}

func isUpperCase(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func isLowerCase(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func toUpper(c byte) byte {
	if isLowerCase(c) {
		c -= 'a'
		c += 'A'
	}
	return c
}

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
	return strings.ReplaceAll(id.(string), "\"", "")
}
