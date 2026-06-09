package clear

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

func NewCmdClear() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear <alerting-integration-id>",
		Short: "Clear an alerting profile's webhooks",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return clearRun(cmd, id)
		},
	}

	return cmd
}

func clearRun(cmd *cobra.Command, alertingProfileID int32) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()
	emptyWebhokList := make([]taikuncore.AlertingWebhookDto, 0)

	response, err := myApiClient.Client.AlertingProfilesAPI.AlertingprofilesAssignWebhooks(ctx, alertingProfileID).AlertingWebhookDto(emptyWebhokList).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return

}
