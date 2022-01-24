package repair

import (
	"fmt"

	"github.com/spf13/cobra"
)

type RepairOptions struct {
	// FIXME add options
}

func NewCmdRepair() *cobra.Command {
	var opts RepairOptions

	cmd := cobra.Command{
		Use:   "repair <project-id>",
		Short: "Repair a project's standalone VMs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return repairRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func repairRun(opts *RepairOptions) (err error) {
	fmt.Println("TODO")
	// FIXME
	return
}
