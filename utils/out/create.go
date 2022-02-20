package out

import (
	"errors"

	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out/fields"
)

func PrintResult(resource interface{}, fields fields.Fields) error {
	if config.OutputOnlyID {
		return printResourceID(resource)
	}

	if config.OutputFormat == config.OutputFormatJson {
		return prettyPrintJson(resource)
	}

	return printApiResponseTable(resource, fields)
}

func printResourceID(resource interface{}) error {
	resourceMap, err := jsonObjectToMap(resource)
	if err != nil {
		return cmderr.ProgramError("printResourceID", err)
	}

	id, found := resourceMap["id"]
	if !found {
		return errors.New("response doesn't contain ID")
	}

	Println(resourceIDToString(id))

	return nil
}

func printApiResponseTable(response interface{}, fields fields.Fields) error {
	if config.AllColumns {
		fields.ShowAll()
	} else if len(config.Columns) != 0 {
		if err := fields.SetVisible(config.Columns); err != nil {
			return err
		}
	}

	resourceMap, err := getApiResponseResourceMap(response)
	if err != nil {
		return cmderr.ProgramError("printApiResponseTable", err)
	}

	tab := newTable()

	for _, field := range fields.VisibleFields() {
		value, _ := getValueFromJsonMap(resourceMap, field.JsonPropertyName())
		tab.AppendRow([]interface{}{
			field.Name(),
			trimCellValue(field.Format(value)),
		})
	}

	renderTable(tab)

	return nil
}

func getApiResponseResourceMap(response interface{}) (resourceMap map[string]interface{}, err error) {
	if resourceMap, err = jsonObjectToMap(response); err == nil {
		if nestedMap, ok := resourceMap[api.ResultField].(map[string]interface{}); ok {
			resourceMap = nestedMap
		} else if nestedMap, ok := resourceMap[api.PayloadField].(map[string]interface{}); ok {
			resourceMap = nestedMap
		}
	}

	return
}
