package cmdutils

import "github.com/spf13/cobra"

type runE func(cmd *cobra.Command, args []string) error

func aggregateRunE(previous runE, newRunE runE) runE {
	if previous == nil {
		return newRunE
	}

	return func(cmd *cobra.Command, args []string) error {
		if err := previous(cmd, args); err != nil {
			return err
		}

		return newRunE(cmd, args)
	}
}
