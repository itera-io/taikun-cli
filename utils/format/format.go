package format

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/itera-io/taikun-cli/config"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

var emptyStruct struct{}

const prettyPrintPrefix = ""
const prettyPrintIndent = "    "

func marshalJsonData(data interface{}) []byte {
	if data == nil {
		data = emptyStruct
	}
	jsonBytes, err := json.MarshalIndent(data, prettyPrintPrefix, prettyPrintIndent)
	if err != nil {
		log.Fatal(err)
	}
	return jsonBytes
}

func PrettyPrintJson(data interface{}) {
	fmt.Println(string(marshalJsonData(data)))
}

func structToMap(data interface{}) map[string]interface{} {
	var m map[string]interface{}
	if err := json.Unmarshal(marshalJsonData(data), &m); err != nil {
		log.Fatal(err)
	}
	return m
}

func structsToMaps(structs []interface{}) []map[string]interface{} {
	maps := make([]map[string]interface{}, len(structs))
	for i, s := range structs {
		maps[i] = structToMap(s)
	}
	return maps
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

func fieldsToHeaderRow(fields []string) table.Row {
	row := make([]interface{}, len(fields))
	for i, field := range fields {
		row[i] = formatFieldName(field)
	}
	return row
}

func newTable() table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleDefault)
	t.Style().Options = table.OptionsNoBorders
	t.Style().Format.Header = text.FormatDefault
	return t
}

const maxColumnWidth = 40
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

func resourceMapToRow(resourceMap map[string]interface{}, fields []string) []interface{} {
	row := make([]interface{}, len(fields))
	for i, field := range fields {
		if value, found := resourceMap[field]; found && value != nil {
			row[i] = trimCellValue(value)
		} else {
			row[i] = ""
		}
	}
	return row
}

func PrettyPrintApiResponseTable(resource interface{}, fields ...string) {
	t := newTable()

	t.AppendHeader(fieldsToHeaderRow(fields))
	t.AppendSeparator()

	resourceMap := structToMap(resource)["result"].(map[string]interface{})
	row := resourceMapToRow(resourceMap, fields)
	t.AppendRow(row)

	t.Render()
}

func PrettyPrintTable(resources interface{}, fields ...string) {
	t := newTable()

	t.AppendHeader(fieldsToHeaderRow(fields))
	t.AppendSeparator()

	resourceMaps := structsToMaps(resources.([]interface{}))
	for _, resourceMap := range resourceMaps {
		t.AppendRow(resourceMapToRow(resourceMap, fields))
	}

	t.Render()
}

func PrintDeleteSuccess(resourceName string, id interface{}) {
	fmt.Printf("%s with ID ", resourceName)
	fmt.Print(id)
	fmt.Println(" was deleted successfully.")
}

func PrintStandardSuccess() {
	fmt.Println("Operation was successful.")
}

func PrintCheckSuccess(name string) {
	fmt.Printf("%s is valid.\n", name)
}
