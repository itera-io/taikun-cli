package out

import (
	"log"

	"github.com/itera-io/taikun-cli/config"
)

// Allows printing of resources of different types in a common table.
// A 'type' column will be added to the table,
// the value of the 'type' cell for the resources
// contained in the slice at index *i* of *resourceSlices*
// will be the type at index *i* of *resourceTypes*.
// Thus, *resourceSlices* and *resourceTypes* MUST have the same length.
func PrintMultipleResults(
	resourceSlices []interface{},
	resourceTypes []string,
	fields ...string,
) {
	if config.OutputFormat == config.OutputFormatJson {
		for _, slice := range resourceSlices {
			prettyPrintJson(slice)
		}
	} else if config.OutputFormat == config.OutputFormatTable {
		if len(resourceSlices) != len(resourceTypes) {
			log.Fatal("PrintMultipleResults: resourcesSlices and resourceTypes must have the same length")
		}

		t := newTable()

		if len(config.Columns) != 0 {
			fields = config.Columns
		}

		fieldsPlusType := append(fields, "type")

		appendHeader(t, fieldsPlusType)

		for resourceIndex, resourcesData := range resourceSlices {
			resources := resourcesData.([]interface{})
			resourceMaps := jsonObjectsToMaps(resources)
			for _, resourceMap := range resourceMaps {
				row := resourceMapToRow(resourceMap, fields)
				row = append(row, resourceTypes[resourceIndex])
				t.AppendRow(row)
			}
		}

		renderTable(t)
	}
}
