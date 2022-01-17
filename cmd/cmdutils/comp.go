package cmdutils

import (
	"log"

	"github.com/spf13/cobra"
)

type completionCoreFunc func(cmd *cobra.Command, args []string, toComplete string) []string

func SetFlagCompletionFunc(cmd *cobra.Command, flagName string, f completionCoreFunc) {
	if err := cmd.RegisterFlagCompletionFunc(flagName, makeCompletionFunc(f)); err != nil {
		log.Fatal(err)
	}
}

func makeCompletionFunc(f completionCoreFunc) func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return f(cmd, args, toComplete), cobra.ShellCompDirectiveNoFileComp
	}
}

func SetFlagCompletionValues(cmd *cobra.Command, flagName string, values ...string) {
	SetFlagCompletionFunc(cmd, flagName, func(cmd *cobra.Command, args []string, toComplete string) []string {
		return values
	})
}
