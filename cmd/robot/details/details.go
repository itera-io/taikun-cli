package details

import (
	"context"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

var detailsFields = fields.New(
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

func NewCmdDetails() *cobra.Command {
	cmd := cobra.Command{
		Use:   "details",
		Short: "Get robot user details",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return detailsRun()
		},
	}

	cmdutils.AddColumnsFlag(&cmd, detailsFields)

	return &cmd
}

func detailsRun() (err error) {
	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.RobotAPI.RobotDetails(context.TODO()).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	return out.PrintResult(data, detailsFields)
}
