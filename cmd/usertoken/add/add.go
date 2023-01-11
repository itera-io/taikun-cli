package add

import (
	"errors"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/usertoken/complete"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/user_token"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

var addFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ACCESS-KEY", "AccessKey",
		),
		field.NewVisible(
			"SECRET-KEY", "SecretKey",
		),
	},
)

type AddOptions struct {
	Name           string
	ExpirationDate string
	ReadOnly       bool
	Endpoints      []string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <command>",
		Short: "Add an user token",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			return addRun(&opts)
		},
		Args: cobra.ExactArgs(1),
	}

	cmd.Flags().StringVar(&opts.ExpirationDate, "expiration-date", "", "The user token expiration date")
	cmd.Flags().BoolVar(&opts.ReadOnly, "is-read-only", false, "Enable to create a user token with read-only permissions")

	cmd.Flags().StringSliceVar(&opts.Endpoints, "endpoints", []string{}, "Endpoints the user token have access to")
	cmdutils.SetFlagCompletionFunc(&cmd, "endpoints", complete.EndpointsCompleteFunc)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	body := &models.UserTokenCreateCommand{
		IsReadonly: opts.ReadOnly,
	}

	if opts.ExpirationDate != "" {
		expiredAt := types.StrToDateTime(opts.ExpirationDate)
		body.ExpireDate = &expiredAt
	}

	if len(opts.Endpoints) != 0 {
		endpoints := []*models.AvailableEndpointData{}
		for i := 0; i < len(opts.Endpoints); i++ {
			endpoint := complete.StringToEndpointFormat(opts.Endpoints[i])
			if endpoint == nil {
				err = errors.New("UserToken: Failed to retrieve endpoint " + opts.Endpoints[i] + ".")
				break
			}
			endpoints = append(endpoints, endpoint)
		}
		body.Endpoints = endpoints
	}
	if err != nil {
		return
	}

	params := user_token.NewUserTokenCreateParams().WithV(taikungoclient.Version).WithBody(body)

	response, err := apiClient.Client.UserToken.UserTokenCreate(params, apiClient)
	if err == nil {
		return out.PrintResult(response.Payload, addFields)
	}

	return
}
