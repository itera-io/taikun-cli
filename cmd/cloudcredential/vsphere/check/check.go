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
	Url      string
	Username string
	Password string
}

func NewCmdCheck() *cobra.Command {
	var opts CheckOptions

	cmd := &cobra.Command{
		Use:   "check",
		Short: "Check the validity of an vSphere cloud credential",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return checkRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Url, "url", "u", "", "vSphere endpoint url (required)")
	cmdutils.MarkFlagRequired(cmd, "url")

	cmd.Flags().StringVarP(&opts.Username, "username", "i", "", "vSphere username (required)")
	cmdutils.MarkFlagRequired(cmd, "username")

	cmd.Flags().StringVarP(&opts.Password, "password", "s", "", "vSphere password (required)")
	cmdutils.MarkFlagRequired(cmd, "password")

	return cmd
}

func checkRun(opts *CheckOptions) (err error) {
	// Create an authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.ValidateVsphereCommand{
		Url:      *taikuncore.NewNullableString(&opts.Url),
		Username: *taikuncore.NewNullableString(&opts.Username),
		Password: *taikuncore.NewNullableString(&opts.Password),
	}

	// Execute a query into the API + graceful exit
	myRequest := myApiClient.Client.VsphereCloudCredentialAPI.VsphereValidate(context.TODO()).ValidateVsphereCommand(body)
	response, err := myRequest.Execute()

	if err == nil {
		out.PrintCheckSuccess("vSphere cloud credential")
	}

	// Did it fail because the request failed (e.g. cannot connect to Taikun) or because the credentials are not valid?
	if err != nil {
		myError := tk.CreateError(response, err)
		myStringError := fmt.Sprint(myError)
		if strings.Contains(myStringError, "Failed to validate") {
			err = cmderr.ErrCheckFailure("vSphere cloud credential") // Taikun responded that credentials are not valid.
		} else {
			err = tk.CreateError(response, err) // Something else happened
		}

		return
	}

	return

}
