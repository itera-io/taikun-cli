package cmdutils

import "github.com/spf13/cobra"

type runE func(cmd *cobra.Command, args []string) error

func aggregateRunE(previous runE, f runE) runE {
	if previous == nil {
		return f
	}
	return func(cmd *cobra.Command, args []string) error {
		if err := previous(cmd, args); err != nil {
			return err
		}
		return f(cmd, args)
	}
}
