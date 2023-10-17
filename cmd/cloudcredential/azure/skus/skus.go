package skus

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/offers"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/publishers"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

type SKUsOptions struct {
	CloudCredentialID int32
	Publisher         string
	Offer             string
	Limit             int32
}

func NewCmdSKUs() *cobra.Command {
	var opts SKUsOptions

	cmd := cobra.Command{
		Use:   "skus <cloud-credential-id>",
		Short: "List the image SKUs of an Azure cloud credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.CloudCredentialID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return skusRun(&opts)
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

	cmd.Flags().StringVarP(&opts.Offer, "offer", "o", "", "Image offer (required)")
	cmdutils.MarkFlagRequired(&cmd, "offer")
	cmdutils.SetFlagCompletionFunc(&cmd, "offer",
		func(cmd *cobra.Command, args []string, toComplete string) (completions []string) {
			completions = make([]string, 0)

			if len(args) >= 1 && opts.Publisher != "" {
				cloudCredentialID, err := types.Atoi32(args[0])
				if err == nil {
					opts := offers.OffersOptions{
						CloudCredentialID: cloudCredentialID,
						Publisher:         opts.Publisher,
					}
					completions, _ = offers.ListOffers(&opts)
				}
			}

			return
		},
	)

	cmdutils.AddLimitFlag(&cmd, &opts.Limit)

	return &cmd
}

func skusRun(opts *SKUsOptions) (err error) {
	skus, err := ListSKUs(opts)
	if err == nil {
		out.PrintStringSlice(skus)
	}

	return
}

func ListSKUs(opts *SKUsOptions) (skus []string, err error) {
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.AzureCloudCredentialAPI.AzureSkus(context.TODO(), opts.CloudCredentialID, opts.Publisher, opts.Offer)
	skus = make([]string, 0)
	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			return nil, tk.CreateError(response, err)
		}

		skus = append(skus, data.Data...)

		count := int32(len(skus))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}

		if count == data.GetTotalCount() {
			break
		}

		myRequest = myRequest.Offset(count)
	}

	if opts.Limit != 0 && int32(len(skus)) > opts.Limit {
		skus = skus[:opts.Limit]
	}

	return skus, nil

	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return nil, err
		}

		params := azure.NewAzureSkusParams().WithV(taikungoclient.Version)
		params = params.WithCloudID(opts.CloudCredentialID)
		params = params.WithPublisher(opts.Publisher)
		params = params.WithOffer(opts.Offer)

		skus = make([]string, 0)

		for {
			response, err := apiClient.Client.Azure.AzureSkus(params, apiClient)
			if err != nil {
				return nil, err
			}

			skus = append(skus, response.Payload.Data...)

			count := int32(len(skus))
			if opts.Limit != 0 && count >= opts.Limit {
				break
			}

			if count == response.Payload.TotalCount {
				break
			}

			params = params.WithOffset(&count)
		}

		if opts.Limit != 0 && int32(len(skus)) > opts.Limit {
			skus = skus[:opts.Limit]
		}

		return skus, nil
	*/
}
