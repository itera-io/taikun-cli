package api

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
var SortDirection = "asc"

// Reverse the sort direction used when listing resources
func ReverseSortDirection() {
	SortDirection = "desc"
}
