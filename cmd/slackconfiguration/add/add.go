package add

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

var addFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
	},
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
	cmdutils.SetFlagCompletionValues(&cmd, "type", types.SlackTypes.Keys()...)

	cmd.Flags().StringVarP(&opts.URL, "url", "u", "", "URL (required)")
	cmdutils.MarkFlagRequired(&cmd, "url")

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")

	cmdutils.AddOutputOnlyIDFlag(&cmd)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.CreateSlackConfigurationCommand{
		Name:      *taikuncore.NewNullableString(&opts.Name),
		Url:       *taikuncore.NewNullableString(&opts.URL),
		Channel:   *taikuncore.NewNullableString(&opts.Channel),
		SlackType: types.GetSlackType(opts.Type),
	}
	if opts.OrganizationID != 0 {
		body.SetOrganizationId(opts.OrganizationID)
	}

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.SlackAPI.SlackCreate(context.TODO()).CreateSlackConfigurationCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	// Manipulate the gathered data
	if err == nil {
		payload := map[string]interface{}{
			"id": data.GetId(),
		}

		return out.PrintResult(payload, addFields)
	}
	return
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		body := models.CreateSlackConfigurationCommand{
			Channel:   opts.Channel,
			Name:      opts.Name,
			SlackType: types.GetSlackType(opts.Type),
			URL:       opts.URL,
		}

		if opts.OrganizationID != 0 {
			body.OrganizationID = opts.OrganizationID
		}

		params := slack.NewSlackCreateParams().WithV(taikungoclient.Version)
		params = params.WithBody(&body)

		response, err := apiClient.Client.Slack.SlackCreate(params, apiClient)
		if err == nil {
			payload := map[string]interface{}{
				"id": response.Payload.ID,
			}

			return out.PrintResult(payload, addFields)
		}

		return
	*/
}
