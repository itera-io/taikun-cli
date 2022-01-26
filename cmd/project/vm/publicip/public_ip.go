package publicip

import (
	"github.com/itera-io/taikun-cli/cmd/project/vm/publicip/disable"
	"github.com/itera-io/taikun-cli/cmd/project/vm/publicip/enable"
	"github.com/spf13/cobra"
)

func NewCmdPublicIp() *cobra.Command {
	cmd := cobra.Command{
		Use:   "public-ip <command>",
		Short: "Toggle a standalone VM's public IP on or off (only for OpenStack)",
	}

	cmd.AddCommand(disable.NewCmdDisable())
	cmd.AddCommand(enable.NewCmdEnable())

	return &cmd
}
