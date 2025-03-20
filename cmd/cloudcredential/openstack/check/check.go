package check

import (
	"context"
	"fmt"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"strings"
)

type CheckOptions struct {
	Username      string
	Password      string
	AppCredId     string
	AppCredSecret string
	URL           string
	Domain        string
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
	cmd.Flags().StringVarP(&opts.Password, "password", "p", "", "OpenStack Password (required)")
	cmd.MarkFlagsRequiredTogether("username", "password")

	cmd.Flags().StringVarP(&opts.AppCredId, "appcredid", "i", "", "OpenStack Application Credential ID")
	cmd.Flags().StringVarP(&opts.AppCredSecret, "appcredsecret", "s", "", "OpenStack Application Credential Secret")
	cmd.MarkFlagsRequiredTogether("appcredid", "appcredsecret")

	cmd.Flags().StringVarP(&opts.Domain, "domain", "d", "", "OpenStack Domain (required)")
	cmd.MarkFlagsRequiredTogether("domain", "username")

	cmd.Flags().StringVar(&opts.URL, "url", "", "OpenStack URL (required)")
	cmdutils.MarkFlagRequired(cmd, "url")

	return cmd
}

func checkRun(opts *CheckOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.CheckOpenstackCommand{
		OpenStackUrl: *taikuncore.NewNullableString(&opts.URL),
	}

	if opts.Username != "" && opts.Password != "" && opts.Domain != "" {
		body.SetOpenStackDomain(opts.Domain)
		body.SetOpenStackUser(opts.Username)
		body.SetOpenStackPassword(opts.Password)
		body.SetApplicationCredEnabled(false)
	}

	if opts.AppCredId != "" && opts.AppCredSecret != "" {
		body.SetOpenStackUser(opts.AppCredId)
		body.SetOpenStackPassword(opts.AppCredSecret)
		body.SetApplicationCredEnabled(true)
	}

	// Execute a query into the API + graceful exit
	myRequest := myApiClient.Client.CheckerAPI.CheckerOpenstack(context.TODO()).CheckOpenstackCommand(body)
	response, err := myRequest.Execute()

	if err == nil {
		out.PrintCheckSuccess("OpenStack cloud credential")
	}

	// Did it fail because the request failed (e.g. cannot connect to Taikun) or because the credentials are not valid?
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

}
