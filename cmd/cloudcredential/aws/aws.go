package aws

import (
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/aws/add"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/aws/check"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/aws/list"
	"github.com/spf13/cobra"
)

func NewCmdAWS() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aws <command>",
		Short: "Manage AWS cloud credentials",
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(check.NewCmdCheck())
	cmd.AddCommand(list.NewCmdList())

	return cmd
}
