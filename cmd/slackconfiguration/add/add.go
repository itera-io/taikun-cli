package add

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/slack"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type AddOptions struct {
	Channel        string
	Name           string
	OrganizationID int32
	Type           string
	URL            string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <name>",
		Short: "Add a Slack configuration",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			if err := cmdutils.CheckFlagValue("type", opts.Type, types.SlackTypes); err != nil {
				return err
			}
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Channel, "channel", "c", "", "Channel (required)")
	cmdutils.MarkFlagRequired(&cmd, "channel")

	cmd.Flags().StringVarP(&opts.Type, "type", "t", "", "Type (required)")
	cmdutils.MarkFlagRequired(&cmd, "type")
	cmdutils.RegisterStaticFlagCompletion(&cmd, "type", types.SlackTypes.Keys()...)

	cmd.Flags().StringVarP(&opts.URL, "url", "u", "", "URL (required)")
	cmdutils.MarkFlagRequired(&cmd, "url")

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")

	cmdutils.AddOutputOnlyIDFlag(&cmd)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.UpsertSlackConfigurationCommand{
		Channel:   opts.Channel,
		Name:      opts.Name,
		SlackType: types.GetSlackType(opts.Type),
		URL:       opts.URL,
	}

	if opts.OrganizationID != 0 {
		body.OrganizationID = opts.OrganizationID
	}

	params := slack.NewSlackCreateParams().WithV(api.Version)
	params = params.WithBody(&body)

	response, err := apiClient.Client.Slack.SlackCreate(params, apiClient)
	if err == nil {
		payload := map[string]interface{}{
			"id": response.Payload,
		}
		out.PrintResult(payload,
			"id",
		)
	}

	return
}
