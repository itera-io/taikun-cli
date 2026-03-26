package list

import (
	"context"

	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible("ID", "id"),
		field.NewVisible("NAME", "name"),
		field.NewVisible("CLAIM", "claimValue"),
	},
)

func NewCmdListGroups() *cobra.Command {
	cmd := cobra.Command{
		Use:   "list <ACCOUNT_ID>",
		Short: "List groups for a specific account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			accountID, err := types.Atoi32(args[0])
			if err != nil {
				return err
			}
			return listGroups(accountID)
		},
	}
	return &cmd
}

func listGroups(accountID int32) (err error) {
	myApiClient := tk.NewClient()
	var groups = make([]taikuncore.GroupListItem, 0)

	req := myApiClient.Client.GroupsAPI.GroupsList(context.TODO())
	req.AccountId(accountID)

	data, response, err := req.Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	groups = append(groups, data.GetData()...)

	return out.PrintResults(groups, listFields)
}
