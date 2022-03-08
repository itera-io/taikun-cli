package google

import (
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/google/add"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/google/check"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/google/list"
	"github.com/spf13/cobra"
)

func NewCmdGoogle() *cobra.Command {
	cmd := cobra.Command{
		Use:     "google <command>",
		Short:   "Manage Google Cloud Platform credentials",
		Aliases: []string{"gcp"},
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(check.NewCmdCheck())
	cmd.AddCommand(list.NewCmdList())

	return &cmd
}
