package main

import (
	"github.com/itera-io/taikun-cli/cmd/root"
	"os"
)

// Goreleaser sets this variable when building binaries
var (
	version = "dev"
)

func main() {
	rootCmd := root.NewCmdRoot()

	// Activate and format a version root flag based on Goreleaser variables.
	rootCmd.Version = version
	rootCmd.SetVersionTemplate("Taikun CLI {{printf \"version %s\" .Version}}\n")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
