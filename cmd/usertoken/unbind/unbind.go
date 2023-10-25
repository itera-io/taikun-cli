package unbind

import (
	"context"
	"fmt"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/usertoken/complete"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/spf13/cobra"
)

type UnbindOptions struct {
	Endpoints []string
	TokenID   string
	UnBindAll bool
}

func NewCmdUnbind() *cobra.Command {
	var opts UnbindOptions
	var err error

	cmd := cobra.Command{
		Use:   "unbind <user-id>",
		Short: "Unbind endpoints to an user token",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.TokenID, err = complete.UserTokenIDFromUserTokenName(args[0])
			if err != nil {
				return err
			}
			return unbindRun(&opts)
		},
	}

	complete.CompleteArgsWithUserTokenName(&cmd)

	cmd.Flags().StringSliceVar(&opts.Endpoints, "endpoints", []string{}, "Endpoints the user token have access to")
	//cmdutils.MarkFlagRequired(&cmd, "endpoints")
	cmdutils.SetFlagCompletionFunc(&cmd, "endpoints", complete.EndpointsCompleteFunc)
	cmd.Flags().BoolVar(&opts.UnBindAll, "unbind-all", false, "Enable to unbind all available endpoints")

	return &cmd

}

func unbindRun(opts *UnbindOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.BindUnbindEndpointToTokenCommand{
		TokenId: *taikuncore.NewNullableString(&opts.TokenID),
	}

	if len(opts.Endpoints) != 0 && opts.UnBindAll {
		return fmt.Errorf("Please specify unbind-all OR enpoints option")
	}

	// Setting user-specified endpoints
	if len(opts.Endpoints) != 0 {
		var endpoints []taikuncore.AvailableEndpointData
		for i := 0; i < len(opts.Endpoints); i++ {
			// Find each endpoint from string
			endpoint, stringToEndpointError := complete.StringToEndpointRemoveFormat(opts.Endpoints[i], opts.TokenID)
			if stringToEndpointError != nil {
				return stringToEndpointError
			}
			// Find Id speciffic
			endpoints = append(endpoints, *endpoint)
		}
		body.Endpoints = endpoints
	}

	if opts.UnBindAll == true {
		// Get all bound endpoints
		allEndpoints, endpointsError := complete.GetAllBindingEndpoints(opts.TokenID, false) // Get all bound endpoints
		if endpointsError != nil {
			return endpointsError
		}
		// Insert them inside the body
		body.Endpoints = allEndpoints
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.UserTokenAPI.UsertokenBindUnbind(context.TODO()).BindUnbindEndpointToTokenCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return

}
