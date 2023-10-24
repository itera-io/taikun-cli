package remove

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <billing-credential-id>...",
		Short: "Delete one or more billing credentials",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ids, err := cmdutils.ArgsToNumericalIDs(args)
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return cmdutils.DeleteMultiple(ids, deleteRun)
		},
		Aliases: cmdutils.DeleteAliases,
	}

	return cmd
}

func deleteRun(billingCredentialID int32) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.OperationCredentialsAPI.OpscredentialsDelete(context.TODO(), billingCredentialID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintDeleteSuccess("Billing credential", billingCredentialID)
	return
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		params := ops_credentials.NewOpsCredentialsDeleteParams().WithV(taikungoclient.Version).WithID(billingCredentialID)

		_, _, err = apiClient.Client.OpsCredentials.OpsCredentialsDelete(params, apiClient)
		if err == nil {
			out.PrintDeleteSuccess("Billing credential", billingCredentialID)
		}

		return
	*/
}
