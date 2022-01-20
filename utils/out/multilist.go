package out

import (
	"log"

	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out/fields"
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
	fields fields.Fields,
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
	fields fields.Fields,
) {
	addTypeColumn := true
	if len(resourceTypes) == 0 {
		addTypeColumn = false
	} else if len(resourceSlices) != len(resourceTypes) {
		log.Fatal("PrintMultipleResults: resourcesSlices and resourceTypes must have the same length")
	}

	t := newTable()

	if config.AllColumns {
		fields.ShowAll()
		addTypeColumn = true
	} else if len(config.Columns) != 0 {
		fields.SetVisible(config.Columns)
		addTypeColumn = false
	}

	header := fields.VisibleNames()
	if addTypeColumn {
		header = append(header, "TYPE")
	}
	appendHeader(t, header)

	for resourceIndex, resourcesData := range resourceSlices {
		if resourceIndex > 0 {
			appendSeparator(t)
		}
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
