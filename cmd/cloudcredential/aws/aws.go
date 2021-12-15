package aws

import (
	"taikun-cli/cmd/cloudcredential/aws/create"
	"taikun-cli/cmd/cloudcredential/aws/list"

	"github.com/spf13/cobra"
)

func NewCmdAWS() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aws <command>",
		Short: "Manage AWS Cloud Credentials",
	}

	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(list.NewCmdList())

	return cmd
}
