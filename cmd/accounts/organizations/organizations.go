package organizations

import (
	"github.com/itera-io/taikun-cli/cmd/accounts/organizations/available"
	"github.com/itera-io/taikun-cli/cmd/accounts/organizations/info"
	"github.com/itera-io/taikun-cli/cmd/accounts/organizations/list"
	"github.com/spf13/cobra"
)

func NewCmdOrganizations() *cobra.Command {
	cmd := cobra.Command{
		Use:   "organizations <command>",
		Short: "Manage account organizations",
	}

	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(info.NewCmdInfo())
	cmd.AddCommand(available.NewCmdListAvailable())

	return &cmd
}
