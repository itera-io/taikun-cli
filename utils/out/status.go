package out

import "github.com/itera-io/taikun-cli/cmd/cmdutils/options"

func PrintStandardSuccess(opts options.Outputter) {
	Println(opts, "Operation was successful.")
}

func PrintDeleteSuccess(opts options.Outputter, resourceName string, id interface{}) {
	Printf(opts, "%s with ID ", resourceName)
	Print(opts, id)
	Println(opts, " was deleted successfully.")
}

func PrintCheckSuccess(opts options.Outputter, name string) {
	Printf(opts, "%s is valid.\n", name)
}
