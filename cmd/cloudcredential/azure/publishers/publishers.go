package publishers

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

type PublishersOptions struct {
	CloudCredentialID int32
	Limit             int32
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

	cmdutils.AddLimitFlag(&cmd, &opts.Limit)

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
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.AzureCloudCredentialAPI.AzurePublishers(context.TODO(), opts.CloudCredentialID)
	publishers = make([]string, 0)
	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			return nil, tk.CreateError(response, err)
		}

		publishers = append(publishers, data.Data...)

		count := int32(len(publishers))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}

		if count == data.GetTotalCount() {
			break
		}

		myRequest = myRequest.Offset(count)
	}

	if opts.Limit != 0 && int32(len(publishers)) > opts.Limit {
		publishers = publishers[:opts.Limit]
	}

	return publishers, nil

	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return nil, err
		}

		params := azure.NewAzurePublishersParams().WithV(taikungoclient.Version)
		params = params.WithCloudID(opts.CloudCredentialID)

		publishers = make([]string, 0)

		for {
			response, err := apiClient.Client.Azure.AzurePublishers(params, apiClient)
			if err != nil {
				return nil, err
			}

			publishers = append(publishers, response.Payload.Data...)

			count := int32(len(publishers))
			if opts.Limit != 0 && count >= opts.Limit {
				break
			}

			if count == response.Payload.TotalCount {
				break
			}

			params = params.WithOffset(&count)
		}

		if opts.Limit != 0 && int32(len(publishers)) > opts.Limit {
			publishers = publishers[:opts.Limit]
		}

		return publishers, nil
	*/
}
