package utils

import (
	"log"

	"github.com/spf13/cobra"
)

const ApiVersion = "1"

var SortDirection = "asc"

func ReverseSortDirection() {
	SortDirection = "desc"
}

func MarkFlagRequired(cmd *cobra.Command, flag string) {
	if err := cmd.MarkFlagRequired(flag); err != nil {
		log.Fatal(err)
	}
}

func RegisterFlagCompletionFunc(cmd *cobra.Command, flagName string, f func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective)) {
	if err := cmd.RegisterFlagCompletionFunc(flagName, f); err != nil {
		log.Fatal(err)
	}
}
