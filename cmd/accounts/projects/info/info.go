package info

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

var infoFields = fields.New(
	[]*field.Field{
		field.NewVisible("ID", "id"),
		field.NewVisible("NAME", "name"),
		field.NewVisible("ORG-ID", "organizationId"),
		field.NewVisible("ORG-NAME", "organizationName"),
		field.NewVisible("CLOUD-TYPE", "cloudType"),
		field.NewVisible("STATUS", "status"),
		field.NewVisible("HEALTH", "health"),
		field.NewVisible("ALERTS", "alertCount"),
		field.NewVisible("K8S-VERSION", "currentKubernetesVersion"),
		field.NewVisible("MODIFIED-AT", "modifiedAt"),
	},
)

type InfoOptions struct {
	AccountID int32
	ProjectID int32
}

func NewCmdInfo() *cobra.Command {
	var opts InfoOptions

	cmd := cobra.Command{
		Use:   "info <ACCOUNT_ID> <PROJECT_ID>",
		Short: "Get detailed information about a project in an account",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.AccountID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			opts.ProjectID, err = types.Atoi32(args[1])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return infoRun(&opts)
		},
	}

	cmdutils.AddColumnsFlag(&cmd, infoFields)

	return &cmd
}

func infoRun(opts *InfoOptions) (err error) {
	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.AccountsAPI.AccountsAccountProjectDetails(context.TODO(), opts.AccountID, opts.ProjectID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	return out.PrintResult(data, infoFields)
}
