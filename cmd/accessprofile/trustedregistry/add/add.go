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
		field.NewVisible(
			"REGISTRY", "registry",
		),
		field.NewVisible(
			"ACCESS-PROFILE", "accessProfileName",
		),
	},
)

type AddOptions struct {
	AccessProfileID int32
	Registry        string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := &cobra.Command{
		Use:   "add <access-profile-id>",
		Short: "Add a trusted registry to an access profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.AccessProfileID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Registry, "registry", "r", "", "Registry name or URL (required)")
	cmdutils.MarkFlagRequired(cmd, "registry")

	cmdutils.AddOutputOnlyIDFlag(cmd)
	cmdutils.AddColumnsFlag(cmd, addFields)

	return cmd
}

func addRun(opts *AddOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.CreateTrustedRegistriesCommand{
		Registry:        *taikuncore.NewNullableString(&opts.Registry),
		AccessProfileId: &opts.AccessProfileID,
	}

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.TrustedRegistriesAPI.TrustedregistriesCreate(context.TODO()).CreateTrustedRegistriesCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	return out.PrintResult(data, addFields)
}
