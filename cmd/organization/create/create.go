package create

import (
	"taikun-cli/api"
	"taikun-cli/utils"

	"github.com/itera-io/taikungoclient/client/organizations"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

func NewCmdCreate() *cobra.Command {
	var opts models.OrganizationCreateCommand

	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create organization",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			opts.IsEligibleUpdateSubscription = true
			return createRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.FullName, "full-name", "f", "", "Full name (required)")
	utils.MarkFlagRequired(cmd, "full-name")

	cmd.Flags().StringVarP(&opts.Address, "address", "a", "", "Address")
	cmd.Flags().StringVarP(&opts.BillingEmail, "billing-email", "b", "", "Billing email")
	cmd.Flags().StringVar(&opts.City, "city", "", "City")
	cmd.Flags().StringVar(&opts.Country, "country", "", "Country")
	cmd.Flags().Float64VarP(&opts.DiscountRate, "discount-rate", "d", 100, "Discount rate")
	cmd.Flags().StringVarP(&opts.Email, "email", "e", "", "Email")
	cmd.Flags().StringVarP(&opts.Phone, "phone", "p", "", "Phone")
	cmd.Flags().StringVarP(&opts.VatNumber, "vat-number", "v", "", "VAT number")

	return cmd
}

func createRun(opts *models.OrganizationCreateCommand) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := organizations.NewOrganizationsCreateParams().WithV(utils.ApiVersion).WithBody(opts)
	_, err = apiClient.Client.Organizations.OrganizationsCreate(params, apiClient)
	if err == nil {
		utils.PrintStandardSuccess()
	}

	return
}
