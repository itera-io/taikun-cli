package remove

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/showbackclient/showback_credentials"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := cobra.Command{
		Use:   "delete <showback-credential-id>...",
		Short: "Delete one or more showback credentials",
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

	return &cmd
}

func deleteRun(showbackCredentialID int32) (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	params := showback_credentials.NewShowbackCredentialsDeleteParams().WithV(taikungoclient.Version).WithID(showbackCredentialID)

	_, err = apiClient.ShowbackClient.ShowbackCredentials.ShowbackCredentialsDelete(params, apiClient)
	if err == nil {
		out.PrintDeleteSuccess("Showback credential", showbackCredentialID)
	}

	return
}
