package apiconfig

const Version = "1"

const DefaultAccessProfileName = "default"
const DefaultAlertingProfileName = "default"
const DefaultKubernetesProfileName = "default"

const ResultField = "result"

var SortDirection = "asc"

func ReverseSortDirection() {
	SortDirection = "desc"
}
