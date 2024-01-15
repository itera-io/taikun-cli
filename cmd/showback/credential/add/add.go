package add

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	tk "github.com/itera-io/taikungoclient"
	taikunshowback "github.com/itera-io/taikungoclient/showbackclient"
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
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikunshowback.CreateShowbackCredentialCommand{
		Name:           *taikunshowback.NewNullableString(&opts.Name),
		Url:            *taikunshowback.NewNullableString(&opts.URL),
		Username:       *taikunshowback.NewNullableString(&opts.Username),
		Password:       *taikunshowback.NewNullableString(&opts.Password),
		OrganizationId: *taikunshowback.NewNullableInt32(&opts.OrganizationID),
	}

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.ShowbackClient.ShowbackCredentialsAPI.ShowbackcredentialsCreate(context.TODO()).CreateShowbackCredentialCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	return out.PrintResult(data, addFields)

}
