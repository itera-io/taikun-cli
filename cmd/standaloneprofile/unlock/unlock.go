package unlock

import "github.com/spf13/cobra"

type UnlockOptions struct {
	// FIXME add options
}

func NewCmdUnlock() *cobra.Command {
	var opts UnlockOptions

	cmd := cobra.Command{
		Use:   "unlock <standalone-profile-id>",
		Short: "Unlock a standalone profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return unlockRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func unlockRun(opts *UnlockOptions) (err error) {
	// FIXME
	return
}
