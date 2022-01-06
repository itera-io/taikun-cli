package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/format"
	"github.com/itera-io/taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/ssh_users"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	AccessProfileID int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list <access-profile-id>",
		Short: "List access profile's SSH users",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			accessProfileID, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.IDArgumentNotANumberError
			}
			opts.AccessProfileID = accessProfileID
			return listRun(&opts)
		},
	}

	cmdutils.AddLimitFlag(cmd)

	return cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := ssh_users.NewSSHUsersListParams().WithV(apiconfig.Version).WithAccessProfileID(opts.AccessProfileID)
	response, err := apiClient.Client.SSHUsers.SSHUsersList(params, apiClient)
	if err != nil {
		return err
	}
	sshUsers := response.Payload

	if config.Limit != 0 && int32(len(sshUsers)) > config.Limit {
		sshUsers = sshUsers[:config.Limit]
	}

	format.PrintResults(sshUsers,
		"id",
		"name",
		"sshPublicKey",
	)
	return
}
