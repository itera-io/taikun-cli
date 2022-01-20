package lock

import "github.com/spf13/cobra"

type LockOptions struct {
	// FIXME add options
}

func NewCmdLock() *cobra.Command {
	var opts LockOptions

	cmd := cobra.Command{
		Use:   "lock <standalone-profile-id>",
		Short: "Lock a standalone profile",
		Args:  cobra.ExactArgs(1), // FIXME maybe
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return lockRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func lockRun(opts *LockOptions) (err error) {
	// FIXME
	return
}
