package out

func PrintStandardSuccess() {
	Println("Operation was successful.")
}

func PrintDeleteSuccess(resourceName string, id interface{}) {
	Printf("%s with ID ", resourceName)
	Print(id)
	Println(" was deleted successfully.")
}

func PrintCheckSuccess(name string) {
	Printf("%s is valid.\n", name)
}
