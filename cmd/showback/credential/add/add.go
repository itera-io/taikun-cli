package add

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"

	"github.com/itera-io/taikungoclient/client/showback"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

var addFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"NAME", "name",
		),
		field.NewVisible(
			"ORG", "organizationName",
		),
		field.NewVisible(
			"ORG-ID", "organizationId",
		),
		field.NewVisible(
			"URL", "url",
		),
		field.NewVisible(
			"USERNAME", "username",
		),
		field.NewHidden(
			"PASSWORD", "password",
		),
		field.NewVisible(
			"LOCK", "isLocked",
		),
		field.NewVisibleWithToStringFunc(
			"CREATED-AT", "createdAt", out.FormatDateTimeString,
		),
		field.NewVisible(
			"CREATED-BY", "createdBy",
		),
	},
)

type AddOptions struct {
	Name           string
	OrganizationID int32
	Password       string
	URL            string
	Username       string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <name>",
		Short: "Add a showback credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Password, "password", "p", "", "Password (Prometheus or other) (required)")
	cmdutils.MarkFlagRequired(&cmd, "password")

	cmd.Flags().StringVarP(&opts.Username, "username", "l", "", "Username (Prometheus or other) (required)")
	cmdutils.MarkFlagRequired(&cmd, "username")

	cmd.Flags().StringVarP(&opts.URL, "url", "u", "", "URL of the source (required)")
	cmdutils.MarkFlagRequired(&cmd, "url")

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddOutputOnlyIDFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, addFields)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.CreateShowbackCredentialCommand{
		Name:           opts.Name,
		OrganizationID: opts.OrganizationID,
		Password:       opts.Password,
		URL:            opts.URL,
		Username:       opts.Username,
	}

	params := showback.NewShowbackCreateCredentialParams().WithV(api.Version)
	params = params.WithBody(&body)

	response, err := apiClient.Client.Showback.ShowbackCreateCredential(params, apiClient)
	if err == nil {
		out.PrintResult(response.Payload, addFields)
	}

	return
}
