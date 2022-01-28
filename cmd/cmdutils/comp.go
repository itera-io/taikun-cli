package cmdutils

import (
	"github.com/spf13/cobra"
)

type CompletionCoreFunc func(cmd *cobra.Command, args []string, toComplete string) []string

func SetFlagCompletionFunc(cmd *cobra.Command, flagName string, f CompletionCoreFunc) {
	if err := cmd.RegisterFlagCompletionFunc(flagName, makeCompletionFunc(f)); err != nil {
		panic(err)
	}
}

func SetArgsCompletionFunc(cmd *cobra.Command, f CompletionCoreFunc) {
	cmd.ValidArgsFunction = makeCompletionFunc(f)
}

func makeCompletionFunc(f CompletionCoreFunc) func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return f(cmd, args, toComplete), cobra.ShellCompDirectiveNoFileComp
	}
}

func SetFlagCompletionValues(cmd *cobra.Command, flagName string, values ...string) {
	SetFlagCompletionFunc(cmd, flagName, func(cmd *cobra.Command, args []string, toComplete string) []string {
		return values
	})
}
