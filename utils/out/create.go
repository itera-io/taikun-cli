package out

import (
	"fmt"
	"os"
	"sort"

	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/config"
)

func PrintResult(resource interface{}, fields ...string) {
	if config.OutputOnlyID {
		printResourceID(resource)
	} else {
		if config.OutputFormat == config.OutputFormatJson {
			prettyPrintJson(resource)
		} else if config.OutputFormat == config.OutputFormatTable {
			printApiResponseTable(resource, fields...)
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

func printApiResponseTable(response interface{}, fields ...string) {
	if len(config.Columns) != 0 {
		fields = config.Columns
	}

	resourceMap := getApiResponseResourceMap(response)
	if len(fields) == 0 {
		printDetailedApiResponseTable(resourceMap)
	}

	t := newTable()
	for _, field := range fields {
		if resourceMap[field] != nil && resourceMap[field] != "" {
			t.AppendRow([]interface{}{formatFieldName(field), resourceMap[field]})
		}
	}

	renderTable(t)
}

func printDetailedApiResponseTable(resourceMap map[string]interface{}) {
	t := newTable()
	keys := make([]string, 0)
	for key, _ := range resourceMap {
		keys = append(keys, key)
	}
	sortedKeys := sort.StringSlice(keys)
	sort.Slice(sortedKeys, sortedKeys.Less)
	for _, key := range sortedKeys {
		if isSimpleApiType(resourceMap[key]) {
			t.AppendRow([]interface{}{formatFieldName(key), resourceMap[key]})
		}
	}
	renderTable(t)
}

func isSimpleApiType(v interface{}) bool {
	if _, simple := v.(string); simple {
		return true
	}
	if _, simple := v.(int32); simple {
		return true
	}
	if _, simple := v.(float64); simple {
		return true
	}
	if _, simple := v.(bool); simple {
		return true
	}
	return false
}

func getApiResponseResourceMap(response interface{}) map[string]interface{} {
	resourceMap := jsonObjectToMap(response)
	if resourceMap[api.ResultField] != nil {
		resourceMap = resourceMap[api.ResultField].(map[string]interface{})
	} else if resourceMap[api.PayloadField] != nil {
		resourceMap = resourceMap[api.PayloadField].(map[string]interface{})
	}
	return resourceMap
}
