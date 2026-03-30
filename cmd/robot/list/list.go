package list

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

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible("USERID", "userId"),
		field.NewVisible("ACCOUNTID", "accountId"),
		field.NewVisible("ACCOUNTNAME", "accountName"),
		field.NewVisible("ACCESSKEY", "accessKey"),
		field.NewVisible("ORGID", "organizationId"),
		field.NewVisible("ORGNAME", "organizationName"),
		field.NewVisible("CREATEDBY", "createdBy"),
		field.NewVisible("NAME", "name"),
		field.NewVisible("DESCRIPTION", "description"),
		field.NewVisible("SCOPES", "scopes"),
		field.NewVisible("IPS", "ips"),
		field.NewVisible("ISACTIVE", "isActive"),
		field.NewVisible("CREATEDAT", "createdAt"),
		field.NewVisible("EXPIRESAT", "expiresAt"),
		field.NewVisible("LASTUSEDAT", "lastUsedAt"),
	},
)

type ListOptions struct {
	OrganizationID int32
	Limit          int32
	Offset         int32
	Search         string
	SearchID       string
}

func NewCmdListRobots() *cobra.Command {
	var opts ListOptions
	cmd := cobra.Command{
		Use:   "list <ACCOUNT_ID>",
		Short: "List robot users for a specific account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			accountID, err := types.Atoi32(args[0])
			if err != nil {
				return err
			}
			return listRobots(accountID, &opts)
		},
	}

	cmdutils.AddOrgIDFlag(&cmd, &opts.OrganizationID)
	cmdutils.AddLimitFlag(&cmd, &opts.Limit)
	cmdutils.AddSortByAndReverseFlags(&cmd, "robot", listFields)
	cmdutils.AddColumnsFlag(&cmd, listFields)
	cmd.Flags().Int32VarP(&opts.Offset, "offset", "", 0, "Offset")
	cmd.Flags().StringVarP(&opts.Search, "search", "", "", "Search")
	cmd.Flags().StringVarP(&opts.SearchID, "search-id", "", "", "Search ID")
	_ = cmd.MarkFlagRequired("organization-id")

	return &cmd
}

func listRobots(accountID int32, opts *ListOptions) (err error) {
	myApiClient := tk.NewClient()
	var robots = make([]taikuncore.RobotUsersListDto, 0)

	req := myApiClient.Client.RobotAPI.RobotList(context.TODO())
	req.AccountId(accountID)
	req.OrganizationId(opts.OrganizationID)
	req.Limit(opts.Limit)
	req.Offset(opts.Offset)
	req.Search(opts.Search)
	req.SearchId(opts.SearchID)

	data, response, err := req.Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	robots = append(robots, data.GetData()...)

	return out.PrintResults(robots, listFields)
}
