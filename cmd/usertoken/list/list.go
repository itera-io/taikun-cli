package list

import (
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/user_token"
	"github.com/itera-io/taikungoclient/models"
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

	cmdutils.AddSortByAndReverseFlags(cmd, "user tokens", ListFields)
	cmdutils.AddColumnsFlag(cmd, ListFields)

	return cmd
}

func listRun(opts *ListOptions) (err error) {
	usertokens, err := ListUserTokens(opts)
	if err != nil {
		return err
	}

	return out.PrintResults(usertokens, ListFields)
}

func ListUserTokens(opts *ListOptions) (userTokenList []*models.UserTokensListDto, err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return nil, err
	}

	params := user_token.NewUserTokenListParams().WithV(taikungoclient.Version)

	userTokenList = make([]*models.UserTokensListDto, 0)

	response, err := apiClient.Client.UserToken.UserTokenList(params, apiClient)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(response.Payload); i++ {
		userTokenList = append(userTokenList, response.Payload[i])
	}

	return userTokenList, nil
}
