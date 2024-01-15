package add

import (
	"context"
	"fmt"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

var addFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"NAME", "name",
		),
		field.NewVisible(
			"FULL-NAME", "fullName",
		),
		field.NewVisible(
			"DISCOUNT-RATE", "discountRate",
		),
		field.NewVisible(
			"PARTNER", "partnerName",
		),
		field.NewVisibleWithToStringFunc(
			"LOCK", "isLocked", out.FormatLockStatus,
		),
		field.NewVisible(
			"READ-ONLY", "isReadOnly",
		),
		field.NewVisible(
			"EMAIL", "email",
		),
		field.NewVisible(
			"BILLING-EMAIL", "billingEmail",
		),
		field.NewVisible(
			"CITY", "city",
		),
		field.NewVisible(
			"COUNTRY", "country",
		),
		field.NewVisible(
			"PHONE", "phone",
		),
		field.NewVisible(
			"VAT", "vatNumber",
		),
		field.NewVisible(
			"SUBSCRIPTION-UPDATES", "isEligibleUpdateSubscription",
		),
		field.NewVisible(
			"PARTNER-ID", "partnerId",
		),
		field.NewHidden(
			"CLOUD-CREDENTIALS", "cloudCredentials",
		),
		field.NewHidden(
			"PROJECTS", "projects",
		),
		field.NewHidden(
			"SERVERS", "servers",
		),
		field.NewHidden(
			"USERS", "users",
		),
		field.NewHiddenWithToStringFunc(
			"CREATED-AT", "createdAt", out.FormatDateTimeString,
		),
	},
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
	cmdutils.SetFlagCompletionFunc(cmd, "country", func(cmd *cobra.Command, args []string, toComplete string) (completions []string) {
		completions = make([]string, 0)

		myApiClient := tk.NewClient()
		data, response, err := myApiClient.Client.CommonAPI.CommonCountries(context.TODO()).Execute()
		if err != nil {
			fmt.Println(fmt.Errorf(tk.CreateError(response, err).Error())) // This function does not return an error... so just call it. #FIXME
			return
		}
		for _, countryListDto := range data {
			completions = append(completions, countryListDto.GetName())
		}

		return

		//apiClient, err := taikungoclient.NewClient()
		//if err != nil {
		//	return
		//}
		//params := common.NewCommonGetCountryListParams().WithV(taikungoclient.Version)
		//result, err := apiClient.Client.Common.CommonGetCountryList(params, apiClient)
		//if err != nil {
		//	return
		//}
		//for _, countryListDto := range result.Payload {
		//	completions = append(completions, countryListDto.Name)
		//}
		//
		//return
	})

	cmdutils.AddOutputOnlyIDFlag(cmd)
	cmdutils.AddColumnsFlag(cmd, addFields)

	return cmd
}

func addRun(opts *AddOptions) (err error) {
	myApiClient := tk.NewClient()
	body := taikuncore.OrganizationCreateCommand{
		Name:                         *taikuncore.NewNullableString(&opts.Name),
		FullName:                     *taikuncore.NewNullableString(&opts.FullName),
		Phone:                        *taikuncore.NewNullableString(&opts.Phone),
		Email:                        *taikuncore.NewNullableString(&opts.Email),
		BillingEmail:                 *taikuncore.NewNullableString(&opts.BillingEmail),
		Address:                      *taikuncore.NewNullableString(&opts.Address),
		Country:                      *taikuncore.NewNullableString(&opts.Country),
		City:                         *taikuncore.NewNullableString(&opts.City),
		VatNumber:                    *taikuncore.NewNullableString(&opts.VatNumber),
		DiscountRate:                 *taikuncore.NewNullableFloat64(&opts.DiscountRate),
		IsEligibleUpdateSubscription: &opts.IsEligibleUpdateSubscription,
	}
	data, response, err := myApiClient.Client.OrganizationsAPI.OrganizationsCreate(context.TODO()).OrganizationCreateCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	return out.PrintResult(data, addFields)

}
