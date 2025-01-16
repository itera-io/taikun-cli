package clear

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
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
			return clearRun(id)
		},
	}

	return cmd
}

func clearRun(alertingProfileID int32) (err error) {
	myApiClient := tk.NewClient()
	emptyWebhokList := make([]taikuncore.AlertingWebhookDto, 0)

	_, response, err := myApiClient.Client.AlertingProfilesAPI.AlertingprofilesAssignWebhooks(context.TODO(), alertingProfileID).AlertingWebhookDto(emptyWebhokList).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return

}
