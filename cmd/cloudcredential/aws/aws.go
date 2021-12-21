package aws

import (
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/aws/check"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/aws/create"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/aws/list"

	"github.com/spf13/cobra"
)

func NewCmdAWS() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aws <command>",
		Short: "Manage AWS Cloud Credentials",
	}

	cmd.AddCommand(check.NewCmdCheck())
	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(list.NewCmdList())

	return cmd
}
