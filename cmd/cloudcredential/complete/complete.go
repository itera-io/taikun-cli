package complete

import (
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/offers"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/publishers"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/skus"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/spf13/cobra"
)

func MakeAzurePublisherCompletionFunc() cmdutils.CompletionCoreFunc {
	return func(cmd *cobra.Command, args []string, toComplete string) (completions []string) {
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
	}
}

func MakeAzureOfferCompletionFunc(publisher *string) cmdutils.CompletionCoreFunc {
	return func(cmd *cobra.Command, args []string, toComplete string) (completions []string) {
		completions = make([]string, 0)

		if len(args) >= 1 && *publisher != "" {
			cloudCredentialID, err := types.Atoi32(args[0])
			if err == nil {
				opts := offers.OffersOptions{
					CloudCredentialID: cloudCredentialID,
					Publisher:         *publisher,
				}
				completions, _ = offers.ListOffers(&opts)
			}
		}

		return
	}
}

func MakeAzureSKUCompletionFunc(publisher *string, offer *string) cmdutils.CompletionCoreFunc {
	return func(cmd *cobra.Command, args []string, toComplete string) (completions []string) {
		completions = make([]string, 0)

		if len(args) >= 1 && *publisher != "" && *offer != "" {
			cloudCredentialID, err := types.Atoi32(args[0])
			if err == nil {
				opts := skus.SKUsOptions{
					CloudCredentialID: cloudCredentialID,
					Publisher:         *publisher,
					Offer:             *offer,
				}
				completions, _ = skus.ListSKUs(&opts)
			}
		}

		return
	}
}
