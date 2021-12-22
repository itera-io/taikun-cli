package format

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/itera-io/taikun-cli/apiconfig"
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
	Println(string(marshalJsonData(data)))
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

	resourceMap := structToMap(resource)
	if resourceMap[apiconfig.ResultField] != nil {
		resourceMap = resourceMap[apiconfig.ResultField].(map[string]interface{})
	}
	row := resourceMapToRow(resourceMap, fields)
	t.AppendRow(row)

	RenderTable(t)
}

func prettyPrintTable(resources []interface{}, fields ...string) {
	t := newTable()

	t.AppendHeader(fieldsToHeaderRow(fields))
	t.AppendSeparator()

	resourceMaps := structsToMaps(resources)
	for _, resourceMap := range resourceMaps {
		t.AppendRow(resourceMapToRow(resourceMap, fields))
	}

	RenderTable(t)
}

func PrintDeleteSuccess(resourceName string, id interface{}) {
	Printf("%s with ID ", resourceName)
	Print(id)
	Println(" was deleted successfully.")
}

func PrintStandardSuccess() {
	Println("Operation was successful.")
}

func PrintCheckSuccess(name string) {
	Printf("%s is valid.\n", name)
}

func trimID(id string) string {
	return strings.ReplaceAll(id, "\"", "")
}

func PrintResourceID(resource interface{}) {
	resourceMap := structToMap(resource)
	if id, found := resourceMap["id"]; found {
		Println(trimID(id.(string)))
	} else {
		fmt.Fprintln(os.Stderr, "ID not found")
	}
}

func PrintResult(resource interface{}, fields ...string) {
	if config.OutputFormat == config.OutputFormatJson {
		PrettyPrintJson(resource)
	} else if config.OutputFormat == config.OutputFormatTable {
		PrettyPrintApiResponseTable(resource, fields...)
	}
}

func interfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func PrintResults(slice interface{}, fields ...string) {
	if config.OutputFormat == config.OutputFormatJson {
		PrettyPrintJson(slice)
	} else if config.OutputFormat == config.OutputFormatTable {
		prettyPrintTable(interfaceSlice(slice), fields...)
	}
}
