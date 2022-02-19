package offers

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/publishers"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/azure"
	"github.com/spf13/cobra"
)

type OffersOptions struct {
	CloudCredentialID int32
	Publisher         string
	Limit             int32
}

func NewCmdOffers() *cobra.Command {
	var opts OffersOptions

	cmd := cobra.Command{
		Use:   "offers <cloud-credential-id>",
		Short: "List the image offers of an Azure cloud credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.CloudCredentialID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return offersRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Publisher, "publisher", "p", "", "Image publisher (required)")
	cmdutils.MarkFlagRequired(&cmd, "publisher")
	cmdutils.SetFlagCompletionFunc(&cmd, "publisher",
		func(cmd *cobra.Command, args []string, toComplete string) (completions []string) {
			completions = make([]string, 0)

			if len(args) >= 1 {
				cloudCredentialID, err := types.Atoi32(args[0])
				if err == nil {
					opts := publishers.PublishersOptions{
						CloudCredentialID: cloudCredentialID,
					}
					completions, _ = publishers.ListPublishers(&opts)
				}
			}

			return
		},
	)

	cmdutils.AddLimitFlag(&cmd, &opts.Limit)

	return &cmd
}

func offersRun(opts *OffersOptions) (err error) {
	offers, err := ListOffers(opts)
	if err == nil {
		out.PrintStringSlice(offers)
	}

	return
}

func ListOffers(opts *OffersOptions) (offers []string, err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return nil, err
	}

	params := azure.NewAzureOffersParams().WithV(api.Version)
	params = params.WithCloudID(opts.CloudCredentialID)
	params = params.WithPublisher(opts.Publisher)

	offers = make([]string, 0)

	for {
		response, err := apiClient.Client.Azure.AzureOffers(params, apiClient)
		if err != nil {
			return nil, err
		}

		offers = append(offers, response.Payload.Data...)

		count := int32(len(offers))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}

		if count == response.Payload.TotalCount {
			break
		}

		params = params.WithOffset(&count)
	}

	if opts.Limit != 0 && int32(len(offers)) > opts.Limit {
		offers = offers[:opts.Limit]
	}

	return offers, nil
}
