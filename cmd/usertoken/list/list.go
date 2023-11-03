package list

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

var ListFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"NAME", "name",
		),
		field.NewVisible(
			"ACCESS-KEY", "accessKey",
		),
		field.NewVisible(
			"READ-ONLY", "isReadonly",
		),
		field.NewVisibleWithToStringFunc(
			"EXPIRATION-DAY", "expireDate", out.FormatDateTimeString,
		),
		field.NewHiddenWithToStringFunc(
			"CREATED", "createdAt", out.FormatDateTimeString,
		),
	},
)

type ListOptions struct {
}

// NewCmdList creates and returns a cobra command for listing user tokens.
func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List user tokens",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Args:    cobra.NoArgs,
		Aliases: cmdutils.ListAliases,
	}

	//cmdutils.AddSortByAndReverseFlags(cmd, "user-tokens", ListFields) // API does not support sorting for this endpoint
	cmdutils.AddColumnsFlag(cmd, ListFields)

	return cmd
}

// listRun calls the API, gets the User Tokens and prints them  in a table.
func listRun(opts *ListOptions) (err error) {
	usertokens, err := ListUserTokens(opts)
	if err != nil {
		return err
	}

	return out.PrintResults(usertokens, ListFields)
}

// ListUserTokens sends a query to the API and returns a list of user tokens.
// Tokens are returned in the UserTokenListDto structs generated in models.
// ListUserTokens is exported because it is used in cmd/usertoken/complete
func ListUserTokens(opts *ListOptions) (userTokenList []*taikuncore.UserTokensListDto, err error) {

	// Connect to the API and retrieve data.
	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.UserTokenAPI.UsertokenList(context.TODO()).Execute()
	if err != nil {
		return nil, tk.CreateError(response, err)
	}

	// Initialise a new, empty slice of UserTokenListDto structs generated in models.
	userTokenList = make([]*taikuncore.UserTokensListDto, 0)

	// For every returned data, create a new line in the fields
	for i := 0; i < len(data); i++ {
		userTokenList = append(userTokenList, &data[i])
	}

	return userTokenList, nil
}
