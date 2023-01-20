package bind

import (
	"errors"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/usertoken/complete"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/user_token"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type BindOptions struct {
	Endpoints []string
	TokenID   string
	BindAll   bool
}

func NewCmdBind() *cobra.Command {
	var opts BindOptions
	var err error

	cmd := cobra.Command{
		Use:   "bind <user-id>",
		Short: "Bind endpoints to an user token",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.TokenID, err = complete.UserTokenIDFromUserTokenName(args[0])
			if err != nil {
				return err
			}
			return bindRun(&opts)
		},
	}

	complete.CompleteArgsWithUserTokenName(&cmd)

	cmd.Flags().StringSliceVar(&opts.Endpoints, "endpoints", []string{}, "Endpoints the user token have access to")
	cmdutils.SetFlagCompletionFunc(&cmd, "endpoints", complete.EndpointsCompleteFunc)
	cmd.Flags().BoolVar(&opts.BindAll, "bind-all", false, "Enable to bind all available endpoints")

	return &cmd

}

func bindRun(opts *BindOptions) (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	body := &models.BindUnbindEndpointToTokenCommand{
		TokenID: opts.TokenID,
		BindAll: opts.BindAll,
	}

	if len(opts.Endpoints) != 0 && opts.BindAll {
		err = errors.New("Please specify bindAll OR enpoints option")
		return
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

	params := user_token.NewUserTokenBindUnbindParams().WithV(taikungoclient.Version).WithBody(body)

	_, err = apiClient.Client.UserToken.UserTokenBindUnbind(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
