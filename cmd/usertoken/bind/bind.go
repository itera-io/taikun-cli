package bind

import (
	"context"
	"fmt"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/usertoken/complete"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
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
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.BindUnbindEndpointToTokenCommand{
		TokenId: *taikuncore.NewNullableString(&opts.TokenID),
		BindAll: &opts.BindAll,
	}

	if len(opts.Endpoints) != 0 && opts.BindAll {
		return fmt.Errorf("Please specify bindAll OR enpoints option")
	}

	// Setting user-specified endpoints
	if len(opts.Endpoints) != 0 && !opts.BindAll {
		var endpoints []taikuncore.AvailableEndpointData
		for i := 0; i < len(opts.Endpoints); i++ {
			// Find each endpoint from string
			endpoint, stringToEndpointError := complete.StringToEndpointBindFormat(opts.Endpoints[i])
			if stringToEndpointError != nil {
				return stringToEndpointError
			}
			endpoints = append(endpoints, *endpoint)
		}
		body.Endpoints = endpoints
		//fmt.Println("-------------------------")
		//fmt.Println(opts.TokenID)
		//fmt.Println(endpoints[0].GetId())
		//fmt.Println(endpoints[0].GetController())
		//fmt.Println(endpoints[0].GetDescription())
		//fmt.Println(endpoints[0].GetMethod())
		//fmt.Println(endpoints[0].GetPath())
		//fmt.Println("-------------------------")
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.UserTokenAPI.UsertokenBindUnbind(context.TODO()).BindUnbindEndpointToTokenCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return
	/*
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
	*/
}
