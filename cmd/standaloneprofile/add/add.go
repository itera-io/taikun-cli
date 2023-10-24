package add

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
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
		field.NewVisibleWithToStringFunc(
			"LOCK", "isLocked", out.FormatLockStatus,
		),
	},
)

type AddOptions struct {
	Name           string
	OrganizationID int32
	PublicKey      string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <name>",
		Short: "Add a standalone profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.PublicKey, "public-key", "p", "", "Public SSH key (required)")
	cmdutils.MarkFlagRequired(&cmd, "public-key")

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddColumnsFlag(&cmd, addFields)
	cmdutils.AddOutputOnlyIDFlag(&cmd)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.StandAloneProfileCreateCommand{
		Name:      *taikuncore.NewNullableString(&opts.Name),
		PublicKey: *taikuncore.NewNullableString(&opts.PublicKey),
	}
	if opts.OrganizationID != 0 {
		body.SetOrganizationId(opts.OrganizationID)
	}

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.StandaloneProfileAPI.StandaloneprofileCreate(context.TODO()).StandAloneProfileCreateCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	// Manipulate the gathered data
	return out.PrintResult(data, addFields)

}
