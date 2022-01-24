package vm

import (
	"github.com/itera-io/taikun-cli/cmd/project/vm/add"
	"github.com/itera-io/taikun-cli/cmd/project/vm/commit"
	"github.com/itera-io/taikun-cli/cmd/project/vm/delete"
	"github.com/itera-io/taikun-cli/cmd/project/vm/disk"
	"github.com/itera-io/taikun-cli/cmd/project/vm/list"
	"github.com/itera-io/taikun-cli/cmd/project/vm/reboot"
	"github.com/itera-io/taikun-cli/cmd/project/vm/repair"
	"github.com/itera-io/taikun-cli/cmd/project/vm/shelve"
	"github.com/itera-io/taikun-cli/cmd/project/vm/start"
	"github.com/itera-io/taikun-cli/cmd/project/vm/tags"
	"github.com/itera-io/taikun-cli/cmd/project/vm/unshelve"
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
	cmd.AddCommand(disk.NewCmdDisk())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(reboot.NewCmdReboot())
	cmd.AddCommand(repair.NewCmdRepair())
	cmd.AddCommand(shelve.NewCmdShelve())
	cmd.AddCommand(start.NewCmdStart())
	cmd.AddCommand(tags.NewCmdTags())
	cmd.AddCommand(unshelve.NewCmdUnshelve())

	return &cmd
}
