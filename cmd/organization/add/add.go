package add

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"

	"github.com/itera-io/taikungoclient/client/common"
	"github.com/itera-io/taikungoclient/client/organizations"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type AddOptions struct {
	Address                      string
	BillingEmail                 string
	City                         string
	Country                      string
	DiscountRate                 float64
	Email                        string
	FullName                     string
	IsEligibleUpdateSubscription bool
	Name                         string
	Phone                        string
	VatNumber                    string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := &cobra.Command{
		Use:   "add <name>",
		Short: "Add an organization",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			opts.IsEligibleUpdateSubscription = true
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.FullName, "full-name", "f", "", "Full name (required)")
	cmdutils.MarkFlagRequired(cmd, "full-name")

	cmd.Flags().StringVarP(&opts.Address, "address", "a", "", "Address")
	cmd.Flags().StringVarP(&opts.BillingEmail, "billing-email", "b", "", "Billing email")
	cmd.Flags().StringVar(&opts.City, "city", "", "City")
	cmd.Flags().Float64VarP(&opts.DiscountRate, "discount-rate", "d", 100, "Discount rate")
	cmd.Flags().StringVarP(&opts.Email, "email", "e", "", "Email")
	cmd.Flags().StringVarP(&opts.Phone, "phone", "p", "", "Phone")
	cmd.Flags().StringVarP(&opts.VatNumber, "vat-number", "v", "", "VAT number")

	cmd.Flags().StringVar(&opts.Country, "country", "", "Country")
	cmdutils.RegisterFlagCompletionFunc(cmd, "country", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		apiClient, err := api.NewClient()
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveDefault
		}

		params := common.NewCommonGetCountryListParams().WithV(api.Version)
		result, err := apiClient.Client.Common.CommonGetCountryList(params, apiClient)
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveDefault
		}

		countryNames := make([]string, 0)
		for _, countryListDto := range result.Payload {
			countryNames = append(countryNames, countryListDto.Name)
		}

		return countryNames, cobra.ShellCompDirectiveDefault
	})

	cmdutils.AddOutputOnlyIDFlag(cmd)

	return cmd
}

func addRun(opts *AddOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.OrganizationCreateCommand{
		Address:                      opts.Address,
		BillingEmail:                 opts.BillingEmail,
		City:                         opts.City,
		Country:                      opts.Country,
		DiscountRate:                 opts.DiscountRate,
		Email:                        opts.Email,
		FullName:                     opts.FullName,
		IsEligibleUpdateSubscription: opts.IsEligibleUpdateSubscription,
		Name:                         opts.Name,
		Phone:                        opts.Phone,
		VatNumber:                    opts.VatNumber,
	}

	params := organizations.NewOrganizationsCreateParams().WithV(api.Version).WithBody(&body)
	response, err := apiClient.Client.Organizations.OrganizationsCreate(params, apiClient)
	if err == nil {
		out.PrintResult(response.Payload,
			"id",
			"name",
			"fullName",
			"discountRate",
			"partnerName",
			"isEligibleUpdateSubscription",
			"isLocked",
			"isReadOnly",
			"users",
			"cloudCredentials",
			"projects",
			"servers",
		)
	}

	return
}
