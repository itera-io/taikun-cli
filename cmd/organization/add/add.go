package add

import (
	"context"
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
	Email    string
	FullName string
	Name     string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := &cobra.Command{
		Use:   "add <name>",
		Short: "Add an organization",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.FullName, "full-name", "f", "", "Full name (required)")
	cmdutils.MarkFlagRequired(cmd, "full-name")

	cmd.Flags().StringVarP(&opts.Email, "email", "e", "", "Email")

	cmdutils.AddOutputOnlyIDFlag(cmd)
	cmdutils.AddColumnsFlag(cmd, addFields)

	return cmd
}

func addRun(opts *AddOptions) (err error) {
	myApiClient := tk.NewClient()
	body := taikuncore.OrganizationCreateCommand{
		Name:     *taikuncore.NewNullableString(&opts.Name),
		FullName: *taikuncore.NewNullableString(&opts.FullName),
		Email:    *taikuncore.NewNullableString(&opts.Email),
	}
	data, response, err := myApiClient.Client.OrganizationsAPI.OrganizationsCreate(context.TODO()).OrganizationCreateCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	return out.PrintResult(data, addFields)

}
