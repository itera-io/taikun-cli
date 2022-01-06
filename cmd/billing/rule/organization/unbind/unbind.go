package unbind

import "github.com/spf13/cobra"

type UnbindOptions struct {
	// FIXME
}

func NewCmdUnbind() *cobra.Command {
	var opts UnbindOptions

	cmd := cobra.Command{
		Use:   "unbind <organization-id>...",
		Short: "Unbind a billing rule from one or more organizations",
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME
			return unbindRun(&opts)
		},
	}

	// FIXME add flags

	return &cmd
}

func unbindRun(opts *UnbindOptions) (err error) {
	// FIXME
	return
}
