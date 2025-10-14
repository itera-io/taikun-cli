package list

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
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

type ListOptions struct {
	AccessProfileID int32
	Limit           int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list <access-profile-id>",
		Short: "List an access profile's trusted registries",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			accessProfileID, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			opts.AccessProfileID = accessProfileID
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmdutils.AddLimitFlag(cmd, &opts.Limit)
	cmdutils.AddColumnsFlag(cmd, listFields)

	return cmd
}

func listRun(opts *ListOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.TrustedRegistriesAPI.TrustedregistriesList(context.TODO(), opts.AccessProfileID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	// Truncate the data
	trustedRegistries := data
	if opts.Limit != 0 && int32(len(trustedRegistries)) > opts.Limit {
		trustedRegistries = trustedRegistries[:opts.Limit]
	}

	return out.PrintResults(trustedRegistries, listFields)
}
