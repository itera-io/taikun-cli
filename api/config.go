package api

import "github.com/itera-io/taikun-cli/config"

// API version
const Version = "1"

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
	var sortDirection string = "asc"
	if config.ReverseSortDirection {
		sortDirection = "desc"
	}

	return &sortDirection
}
