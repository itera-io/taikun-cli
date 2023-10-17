package info

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/organization/list"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

var infoFields = list.ListFields

type InfoOptions struct {
	OrganizationID int32
}

func NewCmdInfo() *cobra.Command {
	var opts InfoOptions

	cmd := cobra.Command{
		Use:   "info <organization-id>",
		Short: "Get detailed information about an organization",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.OrganizationID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			infoFields.ShowAll()
			return infoRun(&opts)
		},
	}

	cmdutils.AddColumnsFlag(&cmd, infoFields)

	return &cmd
}

// infoRun calls the API and gets an object with information which it prints
func infoRun(opts *InfoOptions) (err error) {
	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.OrganizationsAPI.OrganizationsList(context.TODO()).Id(opts.OrganizationID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	if len(data.Data) != 1 {
		return cmderr.ResourceNotFoundError("Organization", opts.OrganizationID)
	}

	return out.PrintResult(data.Data[0], infoFields)
}
