package out

import (
	"fmt"

	"github.com/itera-io/taikun-cli/config"
	"github.com/jedib0t/go-pretty/v6/table"
)

func Println(a ...interface{}) {
	if !config.Quiet {
		fmt.Println(a...)
	}
}

func Printf(format string, a ...interface{}) {
	if !config.Quiet {
		fmt.Printf(format, a...)
	}
}

func Print(a ...interface{}) {
	if !config.Quiet {
		fmt.Print(a...)
	}
}

func renderTable(t table.Writer) {
	if !config.Quiet {
		t.Render()
	}
}

func PrintStringSlice(s []string) {
	if !config.Quiet {
		for _, str := range s {
			fmt.Println(str)
		}
	}
}
