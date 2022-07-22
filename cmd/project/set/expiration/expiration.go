package expiration

import (
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/spf13/cobra"
)

type ExtendLifetimeOptions struct {
	ProjectID int32
	//ExpireAt           date
	DeleteOnExpiration bool
}

func NewCmdExpiration() *cobra.Command {
	var opts ExtendLifetimeOptions

	cmd := cobra.Command{
		Use:   "expiration <project-id>",
		Short: "Manage expiration for a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			//return disableRun(&opts)
			return
		},
	}

	cmd.Flags().BoolVarP(&opts.DeleteOnExpiration, "delete-on-expiration", "del", false, "Delete on expiration (required)")
	cmdutils.MarkFlagRequired(&cmd, "delete-on-expiration")

	// cmd.Flags().StringVarP(&opts.Role, "role", "r", "", "Role (required)")
	// cmdutils.MarkFlagRequired(&cmd, "role")

	return &cmd
}
