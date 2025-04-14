package add

import (
	"context"
	"errors"
	"fmt"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/usertoken/complete"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"time"
)

var addFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ACCESS-KEY", "accessKey",
		),
		field.NewVisible(
			"SECRET-KEY", "secretKey",
		),
	},
)

type AddOptions struct {
	Name           string
	ExpirationDate string
	ReadOnly       bool
	BindAll        bool
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

	cmd.Flags().StringVar(&opts.ExpirationDate, "expiration-date", "", fmt.Sprintf("Expiration date in the format: %s", types.ExpectedDateFormat))
	cmd.Flags().BoolVar(&opts.ReadOnly, "is-read-only", false, "Enable to create a user token with read-only permissions")
	cmd.Flags().BoolVar(&opts.BindAll, "bind-all", false, "Enable to bind all available endpoints")

	cmd.Flags().StringSliceVar(&opts.Endpoints, "endpoints", []string{}, "Endpoints the user token have access to")
	cmdutils.SetFlagCompletionFunc(&cmd, "endpoints", complete.EndpointsCompleteFunc)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	// Prepare body with arguments for the connection
	body := taikuncore.UserTokenCreateCommand{}
	body.SetName(opts.Name)
	body.SetIsReadonly(opts.ReadOnly)
	body.SetBindALL(opts.BindAll)

	// If the expiration date was set by user set it in the request
	if opts.ExpirationDate != "" {
		expiredAt := types.StrToDateTime(opts.ExpirationDate)
		body.SetExpireDate(time.Time(expiredAt))
	}

	// Were some endpoints (or BindAll) set?
	if len(opts.Endpoints) != 0 && opts.BindAll {
		err = errors.New("please specify bindAll OR enpoints option")
		return
	}

	// Setting user-specified endpoints
	if len(opts.Endpoints) != 0 && !opts.BindAll {
		var endpoints []taikuncore.AvailableEndpointData
		for i := 0; i < len(opts.Endpoints); i++ {
			// Find each endpoint from string
			endpoint, stringToEndpointError := complete.StringToEndpointFormat(opts.Endpoints[i], "")
			if stringToEndpointError != nil {
				return stringToEndpointError
			}
			endpoints = append(endpoints, *endpoint)
		}
		body.Endpoints = endpoints
	}

	// Setting all endpoints
	if opts.BindAll {
		// Get all endpoints
		allEndpoints, endpointsError := complete.GetAllEndpoints()
		if endpointsError != nil {
			return endpointsError
		}
		// Insert them inside the body
		body.Endpoints = allEndpoints
	}

	// Send the request to the API and parse the incoming data.
	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.UserTokenAPI.UsertokenCreate(context.TODO()).UserTokenCreateCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	newMap, _ := data.ToMap()
	return out.PrintResult(newMap, addFields)
}
