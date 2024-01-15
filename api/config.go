package api

import "github.com/itera-io/taikun-cli/config"

// Names of default profiles
const DefaultAccessProfileName = "default"
const DefaultAlertingProfileName = "default"
const DefaultKubernetesProfileName = "default"

// Name of field in ApiResponse containing resource info
const ResultField = "result"

// Name of payload field
const PayloadField = "Payload"

// Sort direction to use when listing resources
func GetSortDirection() *string {
	var sortDirection = "asc"
	if config.ReverseSortDirection {
		sortDirection = "desc"
	}

	return &sortDirection
}
