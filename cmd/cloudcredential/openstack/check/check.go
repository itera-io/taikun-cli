package check

import (
	"context"
	"fmt"
	tk "github.com/Smidra/taikungoclient"
	taikuncore "github.com/Smidra/taikungoclient/client"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/spf13/cobra"
	"strings"
)

type CheckOptions struct {
	Username string
	Password string
	URL      string
	Domain   string
}

func NewCmdCheck() *cobra.Command {
	var opts CheckOptions

	cmd := &cobra.Command{
		Use:   "check",
		Short: "Check the validity of an OpenStack cloud credential",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return checkRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Username, "username", "u", "", "OpenStack Username (required)")
	cmdutils.MarkFlagRequired(cmd, "username")

	cmd.Flags().StringVarP(&opts.Password, "password", "p", "", "OpenStack Password (required)")
	cmdutils.MarkFlagRequired(cmd, "password")

	cmd.Flags().StringVarP(&opts.Domain, "domain", "d", "", "OpenStack Domain (required)")
	cmdutils.MarkFlagRequired(cmd, "domain")

	cmd.Flags().StringVar(&opts.URL, "url", "", "OpenStack URL (required)")
	cmdutils.MarkFlagRequired(cmd, "url")

	return cmd
}

func checkRun(opts *CheckOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.CheckOpenstackCommand{
		OpenStackUser:     *taikuncore.NewNullableString(&opts.Username),
		OpenStackPassword: *taikuncore.NewNullableString(&opts.Password),
		OpenStackUrl:      *taikuncore.NewNullableString(&opts.URL),
		OpenStackDomain:   *taikuncore.NewNullableString(&opts.Domain),
	}

	// Execute a query into the API + graceful exit
	myRequest := myApiClient.Client.CheckerAPI.CheckerOpenstack(context.TODO()).CheckOpenstackCommand(body)
	response, err := myRequest.Execute()

	if err == nil {
		out.PrintCheckSuccess("OpenStack cloud credential")
	}

	// Did it fail because the request failed (eg cannot connect to Taikun) or because the credentials are not valid?
	if err != nil {
		myError := tk.CreateError(response, err)
		myStringError := fmt.Sprint(myError)
		if strings.Contains(myStringError, "Failed to validate") {
			err = cmderr.ErrCheckFailure("OpenStack cloud credential") // Taikun responded that credentials are not valid.
		} else {
			err = tk.CreateError(response, err) // Something else happened
		}

		return
	}

	return
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		body := models.CheckOpenstackCommand{
			OpenStackDomain:   opts.Domain,
			OpenStackPassword: opts.Password,
			OpenStackURL:      opts.URL,
			OpenStackUser:     opts.Username,
		}

		params := checker.NewCheckerOpenstackParams().WithV(taikungoclient.Version).WithBody(&body)

		_, err = apiClient.Client.Checker.CheckerOpenstack(params, apiClient)
		if err == nil {
			out.PrintCheckSuccess("OpenStack cloud credential")
		} else if _, isValidationProblem := err.(*checker.CheckerOpenstackBadRequest); isValidationProblem {
			return cmderr.ErrCheckFailure("OpenStack cloud credential")
		}

		return
	*/
}
