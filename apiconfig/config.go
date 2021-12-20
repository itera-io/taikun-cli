package apiconfig

const Version = "1"

const DefaultAccessProfileName = "default"
const DefaultAlertingProfileName = "default"
const DefaultKubernetesProfileName = "default"

var SortDirection = "asc"

func ReverseSortDirection() {
	SortDirection = "desc"
}
