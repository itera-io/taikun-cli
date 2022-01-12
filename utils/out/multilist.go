package out

import (
	"log"

	"github.com/itera-io/taikun-cli/config"
)

// Allows printing of resources of different types into one table.
// If resource types is not empty,
// a 'type' column will be added to the table,
// the value of the 'type' cell for the resources
// contained in the slice at index *i* of *resourceSlices*
// will be the type at index *i* of *resourceTypes*.
// Thus, *resourceSlices* and *resourceTypes* MUST have the same length.
func PrintResultsOfDifferentTypes(
	resourceSlices []interface{},
	resourceTypes []string,
	fields ...string,
) {
	if config.OutputFormat == config.OutputFormatJson {
		for _, slice := range resourceSlices {
			prettyPrintJson(slice)
		}
	} else if config.OutputFormat == config.OutputFormatTable {
		printTableWithDifferentTypes(resourceSlices, resourceTypes, fields)
	}
}

func printTableWithDifferentTypes(
	resourceSlices []interface{},
	resourceTypes []string,
	fields []string,
) {
	addTypeColumn := true
	if len(resourceTypes) == 0 {
		addTypeColumn = false
	} else if len(resourceSlices) != len(resourceTypes) {
		log.Fatal("PrintMultipleResults: resourcesSlices and resourceTypes must have the same length")
	}

	t := newTable()

	if len(config.Columns) != 0 {
		fields = config.Columns
		addTypeColumn = false
	}

	header := fields
	if addTypeColumn {
		header = append(fields, "type")
	}
	appendHeader(t, header)

	for resourceIndex, resourcesData := range resourceSlices {
		resources := resourcesData.([]interface{})
		resourceMaps := jsonObjectsToMaps(resources)
		for _, resourceMap := range resourceMaps {
			row := resourceMapToRow(resourceMap, fields)
			if addTypeColumn {
				row = append(row, resourceTypes[resourceIndex])
			}
			t.AppendRow(row)
		}
	}

	renderTable(t)
}
