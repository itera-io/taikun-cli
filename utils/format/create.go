package format

import (
	"fmt"
	"os"

	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/config"
)

func PrintResult(resource interface{}, fields ...string) {
	if config.OutputOnlyID {
		printResourceID(resource)
	} else {
		if config.OutputFormat == config.OutputFormatJson {
			prettyPrintJson(resource)
		} else if config.OutputFormat == config.OutputFormatTable {
			PrettyPrintApiResponseTable(resource, fields...)
		}
	}
}

func PrintResultVertical(resource interface{}, fields ...string) {
	if config.OutputOnlyID {
		printResourceID(resource)
	} else {
		if config.OutputFormat == config.OutputFormatJson {
			prettyPrintJson(resource)
		} else if config.OutputFormat == config.OutputFormatTable {
			PrettyPrintApiResponseVerticalTable(resource, fields...)
		}
	}
}

func printResourceID(resource interface{}) {
	resourceMap := jsonObjectToMap(resource)
	if id, found := resourceMap["id"]; found {
		Println(resourceIDToString(id))
	} else {
		fmt.Fprintln(os.Stderr, "ID not found")
	}
}

func PrettyPrintApiResponseTable(response interface{}, fields ...string) {
	t := newTable()

	if len(config.Columns) != 0 {
		fields = config.Columns
	}

	resourceMap := getApiResponseResourceMap(response)
	nonEmptyFields := make([]string, 0)
	for _, field := range fields {
		if _, fieldExists := resourceMap[field]; fieldExists {
			nonEmptyFields = append(nonEmptyFields, field)
		}
	}

	if len(nonEmptyFields) == 0 {
		Println("No data")
	} else {
		appendHeader(t, nonEmptyFields)
		row := resourceMapToRow(resourceMap, nonEmptyFields)
		t.AppendRow(row)
		renderTable(t)
	}
}

func PrettyPrintApiResponseVerticalTable(response interface{}, fields ...string) {
	t := newTable()
	appendHeader(t, []string{"field", "value"})

	if len(config.Columns) != 0 {
		fields = config.Columns
	}

	resourceMap := getApiResponseResourceMap(response)
	for _, field := range fields {
		if resourceMap[field] != nil && resourceMap[field] != "" {
			t.AppendRow([]interface{}{formatFieldName(field), resourceMap[field]})
		}
	}

	renderTable(t)
}

func getApiResponseResourceMap(response interface{}) map[string]interface{} {
	resourceMap := jsonObjectToMap(response)
	if resourceMap[apiconfig.ResultField] != nil {
		resourceMap = resourceMap[apiconfig.ResultField].(map[string]interface{})
	} else if resourceMap[apiconfig.PayloadField] != nil {
		resourceMap = resourceMap[apiconfig.PayloadField].(map[string]interface{})
	}
	return resourceMap
}
