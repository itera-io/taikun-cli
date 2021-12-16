package delete

import (
	"taikun-cli/api"
	"taikun-cli/apiconfig"
	"taikun-cli/utils/format"
	"taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/ops_credentials"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <billing-credential-id>",
		Short: "Delete a billing credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.Atoi32(args[0])
			if err != nil {
				return format.WrongIDArgumentFormatError
			}
			return deleteRun(id)
		},
	}

	return cmd
}

func deleteRun(id int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := ops_credentials.NewOpsCredentialsDeleteParams().WithV(apiconfig.Version).WithID(id)
	_, _, err = apiClient.Client.OpsCredentials.OpsCredentialsDelete(params, apiClient)
	if err == nil {
		format.PrintDeleteSuccess("Billing credential", id)
	}

	return
}
