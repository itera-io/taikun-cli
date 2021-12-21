package main

import (
	"os"

	"github.com/itera-io/taikun-cli/cmd/root"
)

func main() {
	rootCmd := root.NewCmdRoot()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
