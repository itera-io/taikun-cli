package main

import (
	"os"
	"taikun-cli/cmd/root"
)

func main() {
	rootCmd := root.NewCmdRoot()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
