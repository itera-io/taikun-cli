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
	Url          string
	ClientId     string
	ClientSecret string
}

func NewCmdCheck() *cobra.Command {
	var opts CheckOptions

	cmd := &cobra.Command{
		Use:   "check",
		Short: "Check the validity of an Proxmox cloud credential",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return checkRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Url, "url", "u", "", "Proxmox endpoint url (required)")
	cmdutils.MarkFlagRequired(cmd, "url")

	cmd.Flags().StringVarP(&opts.ClientId, "token-id", "i", "", "Proxmox Client ID (required)")
	cmdutils.MarkFlagRequired(cmd, "token-id")

	cmd.Flags().StringVarP(&opts.ClientSecret, "token-secret", "s", "", "Proxmox Client secret (required)")
	cmdutils.MarkFlagRequired(cmd, "token-secret")

	return cmd
}

func checkRun(opts *CheckOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.ProxmoxCheckerCommand{
		Url:         *taikuncore.NewNullableString(&opts.Url),
		TokenId:     *taikuncore.NewNullableString(&opts.ClientId),
		TokenSecret: *taikuncore.NewNullableString(&opts.ClientSecret),
	}

	// Execute a query into the API + graceful exit
	myRequest := myApiClient.Client.CheckerAPI.CheckerProxmox(context.TODO()).ProxmoxCheckerCommand(body)
	_, response, err := myRequest.Execute()

	if err == nil {
		out.PrintCheckSuccess("Proxmox cloud credential")
	}

	// Did it fail because the request failed (e.g. cannot connect to Taikun) or because the credentials are not valid?
	if err != nil {
		myError := tk.CreateError(response, err)
		myStringError := fmt.Sprint(myError)
		if strings.Contains(myStringError, "Failed to validate") {
			err = cmderr.ErrCheckFailure("Proxmox cloud credential") // Taikun responded that credentials are not valid.
		} else {
			err = tk.CreateError(response, err) // Something else happened
		}

		return
	}

	return

}
