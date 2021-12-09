package main

import (
	"fmt"
	"os"
	"taikun-cli/api"
	"taikun-cli/cmd/root"
)

func main() {
	apiClient, err := api.NewClient()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	rootCmd := root.NewCmdRoot(apiClient)
	rootCmd.Execute()
}
