package main

import (
	"github.com/itera-io/taikun-cli/cmd/root"
	"os"
)

func main() {
	rootCmd := root.NewCmdRoot()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
