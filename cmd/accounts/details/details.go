package details

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

var detailsFields = fields.New(
	[]*field.Field{
		field.NewVisible("ID", "id"),
		field.NewVisible("NAME", "name"),
		field.NewVisible("EMAIL", "email"),
		field.NewVisible("ORGS", "organizationsCount"),
		field.NewVisible("USERS", "usersCount"),
		field.NewVisible("GROUPS", "groupsCount"),
		field.NewVisible("PROJECTS", "projectsCount"),
		field.NewVisible("CREATED-AT", "createdAt"),
		field.NewVisible("2FA", "is2FAEnabled"),
	},
)

type DetailsOptions struct {
	AccountID int32
}

func NewCmdDetails() *cobra.Command {
	var opts DetailsOptions

	cmd := cobra.Command{
		Use:   "details <ACCOUNT_ID>",
		Short: "Get account details",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.AccountID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return detailsRun(&opts)
		},
	}

	cmdutils.AddColumnsFlag(&cmd, detailsFields)

	return &cmd
}

func detailsRun(opts *DetailsOptions) (err error) {
	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.AccountsAPI.AccountsDetails(context.TODO(), opts.AccountID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	return out.PrintResult(data, detailsFields)
}
