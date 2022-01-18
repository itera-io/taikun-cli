package options

type ListSorter interface {
	GetSortByOption() *string
	GetReverseSortDirectionOption() *bool
}

func GetSortDirection(listSorter ListSorter) *string {
	sortDirection := "asc"
	if *listSorter.GetReverseSortDirectionOption() {
		sortDirection = "desc"
	}
	return &sortDirection
}

type ListLimiter interface {
	GetLimitOption() *int32
}
