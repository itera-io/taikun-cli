package info

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/organization/list"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/organizations"
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

func infoRun(opts *InfoOptions) (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	params := organizations.NewOrganizationsListParams().WithV(taikungoclient.Version)
	params = params.WithID(&opts.OrganizationID)

	response, err := apiClient.Client.Organizations.OrganizationsList(params, apiClient)
	if err != nil {
		return
	}

	if len(response.Payload.Data) != 1 {
		return cmderr.ResourceNotFoundError("Organization", opts.OrganizationID)
	}

	return out.PrintResult(response.Payload.Data[0], infoFields)
}
