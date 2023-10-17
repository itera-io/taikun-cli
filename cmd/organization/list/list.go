package list

import (
	"context"
	tk "github.com/Smidra/taikungoclient"
	taikuncore "github.com/Smidra/taikungoclient/client"
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/spf13/cobra"
)

var ListFields = fields.New(
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
		field.NewHidden(
			"EMAIL", "email",
		),
		field.NewHidden(
			"BILLING-EMAIL", "billingEmail",
		),
		field.NewHidden(
			"CITY", "city",
		),
		field.NewHidden(
			"COUNTRY", "country",
		),
		field.NewHidden(
			"PHONE", "phone",
		),
		field.NewHidden(
			"VAT", "vatNumber",
		),
		field.NewHidden(
			"SUBSCRIPTION-UPDATES", "isEligibleUpdateSubscription",
		),
		field.NewHidden(
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
		field.NewVisibleWithToStringFunc(
			"LOCK", "isLocked", out.FormatLockStatus,
		),
		field.NewHidden(
			"READ-ONLY", "isReadOnly",
		),
		field.NewHiddenWithToStringFunc(
			"CREATED-AT", "createdAt", out.FormatDateTimeString,
		),
	},
)

type ListOptions struct {
	Limit int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list",
		Short: "List organizations",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Args:    cobra.NoArgs,
		Aliases: cmdutils.ListAliases,
	}

	cmdutils.AddSortByAndReverseFlags(&cmd, "organizations", ListFields)
	cmdutils.AddColumnsFlag(&cmd, ListFields)
	cmdutils.AddLimitFlag(&cmd, &opts.Limit)

	return &cmd
}

// listRun sends multiple queries to the API and returns a list of organizations.
// Organizations are returned in the UserForListDto structs generated in models.
func listRun(opts *ListOptions) (err error) {
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.OrganizationsAPI.OrganizationsList(context.TODO())
	// Set Sorting if set in command line options
	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}

	var organizations = make([]taikuncore.OrganizationDetailsDto, 0)

	// Execute the request, it returns only 50 lines in one page
	// then execute it again with an Offset until you have read all of it.
	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			return tk.CreateError(response, err)
		}
		organizations = append(organizations, data.Data...)

		organizationsCount := int32(len(organizations))
		if opts.Limit != 0 && organizationsCount >= opts.Limit {
			break
		}

		if organizationsCount == data.GetTotalCount() {
			break
		}

		myRequest = myRequest.Offset(organizationsCount)
	}

	// We have (over)reached the limit, cut it at the limit and break
	if opts.Limit != 0 && int32(len(organizations)) > opts.Limit {
		organizations = organizations[:opts.Limit]
	}

	return out.PrintResults(organizations, ListFields)
}
