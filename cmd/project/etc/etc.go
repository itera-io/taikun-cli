package etc

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/notifications"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

var etcFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"OPERATION", "operation",
		),
		field.NewVisibleWithToStringFunc(
			"ESTIMATED-TIME", "estimatedTime", out.FormatETC,
		),
		field.NewHidden(
			"PROJECT-ID", "projectId",
		),
	},
)

type EtcOptions struct {
	ProjectID int32
}

func NewCmdEtc() *cobra.Command {
	var opts EtcOptions

	cmd := cobra.Command{
		Use:   "etc <project-id>",
		Short: "Get estimated time of completion for project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return etcRun(&opts)
		},
	}

	cmdutils.AddColumnsFlag(&cmd, etcFields)

	return &cmd
}

func etcRun(opts *EtcOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.GetProjectOperationCommand{
		ProjectID: opts.ProjectID,
	}

	params := notifications.NewNotificationsGetProjectOperationMessagesParams().WithV(api.Version)
	params = params.WithBody(&body)

	response, err := apiClient.Client.Notifications.NotificationsGetProjectOperationMessages(params, apiClient)
	if err == nil {
		out.PrintResult(response, etcFields)
	}

	return
}
