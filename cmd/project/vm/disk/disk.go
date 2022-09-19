package disk

import (
	"github.com/itera-io/taikun-cli/cmd/project/vm/disk/add"
	"github.com/itera-io/taikun-cli/cmd/project/vm/disk/list"
	"github.com/itera-io/taikun-cli/cmd/project/vm/disk/remove"
	"github.com/itera-io/taikun-cli/cmd/project/vm/disk/resize"
	"github.com/spf13/cobra"
)

func NewCmdDisk() *cobra.Command {
	cmd := cobra.Command{
		Use:   "disk <command>",
		Short: "Manage a standalone VM's disks",
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(remove.NewCmdDelete())
	cmd.AddCommand(resize.NewCmdResize())

	return &cmd
}
