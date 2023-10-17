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
		err = errors.New("Please specify bindAll OR enpoints option")
		return
	}

	// Preparing to set user-specified endpoints
	if len(opts.Endpoints) != 0 && !opts.BindAll {
		fmt.Println("Setting up some endpoints...")
		endpoints := []taikuncore.AvailableEndpointData{}
		//endpoints := []*models.AvailableEndpointData{}
		for i := 0; i < len(opts.Endpoints); i++ {
			endpoint := *complete.StringToEndpointFormat(opts.Endpoints[i])
			if _, ok := endpoint.GetIdOk(); ok {
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

	// Set bind all

	// Send the request to the API and parse the incoming data.
	myApiClient := tk.NewClient()
	data, _, err := myApiClient.Client.UserTokenAPI.UsertokenCreate(context.TODO()).UserTokenCreateCommand(body).Execute()
	if err == nil {
		newMap, _ := data.ToMap()
		return out.PrintResult(newMap, addFields)
	}

	return
}
