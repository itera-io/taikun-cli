package main

import (
	"taikun-cli/cmd/root"
)

func main() {
	rootCmd := root.NewCmdRoot()
	rootCmd.Execute()
}
