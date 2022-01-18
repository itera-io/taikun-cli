package out

import (
	"fmt"

	"github.com/itera-io/taikun-cli/cmd/cmdutils/options"
	"github.com/jedib0t/go-pretty/v6/table"
)

func Println(opts options.Outputter, a ...interface{}) {
	if !*opts.GetQuietOption() {
		fmt.Println(a...)
	}
}

func Printf(opts options.Outputter, format string, a ...interface{}) {
	if !*opts.GetQuietOption() {
		fmt.Printf(format, a...)
	}
}

func Print(opts options.Outputter, a ...interface{}) {
	if !*opts.GetQuietOption() {
		fmt.Print(a...)
	}
}

func renderTable(opts options.Outputter, t table.Writer) {
	if !*opts.GetQuietOption() {
		t.Render()
	}
}
