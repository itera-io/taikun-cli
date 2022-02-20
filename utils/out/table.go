package out

import (
	"errors"
	"os"
	"reflect"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func printTable(data interface{}, fields fields.Fields) error {
	tab := newTable()

	if config.AllColumns {
		fields.ShowAll()
	} else if len(config.Columns) != 0 {
		if err := fields.SetVisible(config.Columns); err != nil {
			return err
		}
	}

	appendHeader(tab, fields.VisibleNames())

	resources, err := interfaceToInterfaceSlice(data)
	if err != nil {
		return cmderr.ProgramError("printTable", err)
	}

	if parentObjectName, nested := fields.AreNested(); nested {
		allNestedResources := make([]interface{}, 0)

		for i := range resources {
			nestedResources, err := getNestedResources(resources[i], parentObjectName)
			if err != nil {
				return err
			}

			allNestedResources = append(allNestedResources, nestedResources...)
		}

		resources = allNestedResources
	}

	resourceMaps, err := jsonObjectsToMaps(resources)
	if err != nil {
		return cmderr.ProgramError("printTable", err)
	}

	for _, resourceMap := range resourceMaps {
		tab.AppendRow(resourceMapToRow(resourceMap, fields))
	}

	renderTable(tab)

	return nil
}

func newTable() table.Writer {
	tab := table.NewWriter()

	tab.SetOutputMirror(os.Stdout)
	tab.SetStyle(table.StyleDefault)

	tab.Style().Format.Header = text.FormatDefault
	tab.Style().Options = table.OptionsNoBorders

	if config.NoDecorate {
		tab.Style().Options = table.OptionsNoBordersAndSeparators
		tab.Style().Box.PaddingLeft = ""
	}

	return tab
}

func getNestedResources(resource interface{}, parentObjectName string) (nestedResources []interface{}, err error) {
	resourceMap, err := jsonObjectToMap(resource)
	if err != nil {
		return nil, cmderr.ProgramError("getNestedResource", err)
	}

	nestedData, nestedDataOk := resourceMap[parentObjectName]
	if !nestedDataOk {
		return nil, cmderr.ProgramError("getNestedResource", errors.New("could not find nested resource"))
	}

	nestedResources, nestedResourcesOk := nestedData.([]interface{})
	if !nestedResourcesOk {
		return nil, cmderr.ProgramError("getNestedResource", errors.New("nested resource is not array"))
	}

	return
}

func resourceMapToRow(resourceMap map[string]interface{}, fields fields.Fields) []interface{} {
	row := make([]interface{}, fields.VisibleSize())

	for fieldIndex, field := range fields.VisibleFields() {
		if value, found := getValueFromJsonMap(resourceMap, field.JsonPropertyName()); found && value != nil {
			row[fieldIndex] = field.Format(value)
		} else {
			row[fieldIndex] = ""
		}

		row[fieldIndex] = trimCellValue(row[fieldIndex])
	}

	return row
}

func interfaceToInterfaceSlice(v interface{}) ([]interface{}, error) {
	slice := reflect.ValueOf(v)
	if slice.Kind() != reflect.Slice {
		return nil, errors.New("failed to convert interface to interface slice")
	}

	// Keep the distinction between nil and empty slice input
	if slice.IsNil() {
		return nil, nil
	}

	ret := make([]interface{}, slice.Len())

	for i := 0; i < slice.Len(); i++ {
		ret[i] = slice.Index(i).Interface()
	}

	return ret, nil
}

func stringSliceToRow(fields []string) table.Row {
	row := make([]interface{}, len(fields))
	for i, field := range fields {
		row[i] = field
	}

	return row
}

func appendHeader(t table.Writer, fields []string) {
	if !config.NoDecorate {
		t.AppendHeader(stringSliceToRow(fields))
	}
}

func appendSeparator(t table.Writer) {
	if !config.NoDecorate {
		t.AppendSeparator()
	}
}
