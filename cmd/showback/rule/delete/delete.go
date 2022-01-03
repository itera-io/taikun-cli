package delete

import (
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := cobra.Command{
		Use:   "delete <id>...",
		Short: "Delete one or more showback rules",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmdutils.DeleteMultiple(args, deleteRun)
		},
	}

	return &cmd
}

func deleteRun(id int32) (err error) {
	// FIXME
	return
}
