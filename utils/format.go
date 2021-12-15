package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

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

func PrettyPrintTable(resources interface{}, fields ...string) {
	t := newTable()

	t.AppendHeader(fieldsToHeaderRow(fields))
	t.AppendSeparator()

	resourceMaps := structsToMaps(resources.([]interface{}))
	for _, resourceMap := range resourceMaps {
		row := make([]interface{}, len(fields))
		for i, field := range fields {
			row[i] = resourceMap[field]
		}
		t.AppendRow(table.Row(row))
	}

	t.Render()
}

func PrintDeleteSuccess(resourceName string, id interface{}) {
	fmt.Printf("%s with ID ", resourceName)
	fmt.Print(id)
	fmt.Println(" was deleted successfully.")
}
