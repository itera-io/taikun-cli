package format

import (
	"fmt"

	"github.com/itera-io/taikun-cli/config"
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
