package remove

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/stand_alone_profile"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := cobra.Command{
		Use:   "delete <standalone-profile-id>...",
		Short: "Delete one or more standalone profiles",
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

func deleteRun(standaloneProfileID int32) (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	body := models.DeleteStandAloneProfileCommand{ID: standaloneProfileID}
	params := stand_alone_profile.NewStandAloneProfileDeleteParams().WithV(taikungoclient.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.StandAloneProfile.StandAloneProfileDelete(params, apiClient)
	if err == nil {
		out.PrintDeleteSuccess("Standalone profile", standaloneProfileID)
	}

	return
}
