package list

import (
	"fmt"
	"taikun-cli/api"
	"taikun-cli/config"
	"taikun-cli/utils"

	"github.com/itera-io/taikungoclient/client/ssh_users"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	AccessProfileID int32
	Limit           int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list <access-profile-id>",
		Short: "List access profile's SSH users",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			accessProfileID, err := utils.Atoi32(args[0])
			if err != nil {
				return fmt.Errorf("the given ID must be a number")
			}
			if opts.Limit < 0 {
				return fmt.Errorf("limit flag must be positive")
			}
			if !config.OutputFormatIsValid() {
				return config.OutputFormatInvalidError
			}
			opts.AccessProfileID = accessProfileID
			return listRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.Limit, "limit", "l", 0, "Limit number of results (limitless by default)")

	return cmd
}

func printResults(sshUsers []*models.SSHUsersListDto) {
	if config.OutputFormat == config.OutputFormatJson {
		utils.PrettyPrintJson(sshUsers)
	} else if config.OutputFormat == config.OutputFormatTable {
		data := make([]interface{}, len(sshUsers))
		for i, sshUser := range sshUsers {
			data[i] = sshUser
		}
		utils.PrettyPrintTable(data,
			"id",
			"name",
			"sshPublicKey",
		)
	}
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := ssh_users.NewSSHUsersListParams().WithV(utils.ApiVersion).WithAccessProfileID(opts.AccessProfileID)
	response, err := apiClient.Client.SSHUsers.SSHUsersList(params, apiClient)
	if err != nil {
		return err
	}
	sshUsers := response.Payload

	if opts.Limit != 0 && int32(len(sshUsers)) > opts.Limit {
		sshUsers = sshUsers[:opts.Limit]
	}

	printResults(sshUsers)
	return
}
