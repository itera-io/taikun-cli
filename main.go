package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/itera-io/taikun-cli/cmd/root"
)

// Goreleaser sets this variable when building binaries
var (
	version = "dev"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	rootCmd := root.NewCmdRoot()
	rootCmd.SetContext(ctx)

	// Activate and format a version root flag based on Goreleaser variables.
	rootCmd.Version = version
	rootCmd.SetVersionTemplate("Taikun CLI {{printf \"version %s\" .Version}}\n")

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
