package publishers

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/azure"
	"github.com/spf13/cobra"
)

type PublishersOptions struct {
	CloudCredentialID int32
}

func NewCmdPublishers() *cobra.Command {
	var opts PublishersOptions

	cmd := cobra.Command{
		Use:   "publishers <cloud-credential-id>",
		Short: "List the image publishers of an Azure cloud credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.CloudCredentialID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return publishersRun(&opts)
		},
	}

	return &cmd
}

func publishersRun(opts *PublishersOptions) (err error) {
	publishers, err := ListPublishers(opts)
	if err == nil {
		out.PrintStringSlice(publishers)
	}
	return
}

func ListPublishers(opts *PublishersOptions) (publishers []string, err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := azure.NewAzurePublishersParams().WithV(api.Version)
	params = params.WithCloudID(opts.CloudCredentialID)

	publishers = make([]string, 0)
	for {
		response, err := apiClient.Client.Azure.AzurePublishers(params, apiClient)
		if err != nil {
			return nil, err
		}
		publishers = append(publishers, response.Payload.Data...)
		count := int32(len(publishers))
		if count == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&count)
	}

	return
}
