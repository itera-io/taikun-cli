package spot

import (
	"github.com/itera-io/taikun-cli/cmd/project/spot/full"
	"github.com/itera-io/taikun-cli/cmd/project/spot/vms"
	"github.com/itera-io/taikun-cli/cmd/project/spot/worker"
	"github.com/spf13/cobra"
)

func NewCmdSpot() *cobra.Command {
	cmd := cobra.Command{
		Use:   "spot <command>",
		Short: "Manage if projects can use spot flavors",
	}

	cmd.AddCommand(full.NewCmdFull())
	cmd.AddCommand(worker.NewCmdWorker())
	cmd.AddCommand(vms.NewCmdVms())

	return &cmd
}
