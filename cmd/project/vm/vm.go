package vm

import (
	"github.com/itera-io/taikun-cli/cmd/project/vm/add"
	"github.com/itera-io/taikun-cli/cmd/project/vm/commit"
	"github.com/itera-io/taikun-cli/cmd/project/vm/delete"
	"github.com/itera-io/taikun-cli/cmd/project/vm/list"
	"github.com/itera-io/taikun-cli/cmd/project/vm/repair"
	"github.com/spf13/cobra"
)

func NewCmdVm() *cobra.Command {
	cmd := cobra.Command{
		Use:   "vm <command>",
		Short: "Manage a project's standalone VMs",
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(commit.NewCmdCommit())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(repair.NewCmdRepair())

	return &cmd
}
