package check

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"

	"github.com/itera-io/taikungoclient/client/checker"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
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
		Use:   "check <name>",
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
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.CheckOpenstackCommand{
		OpenStackDomain:   opts.Domain,
		OpenStackPassword: opts.Password,
		OpenStackURL:      opts.URL,
		OpenStackUser:     opts.Username,
	}

	params := checker.NewCheckerOpenstackParams().WithV(api.Version).WithBody(&body)
	_, err = apiClient.Client.Checker.CheckerOpenstack(params, apiClient)
	if err == nil {
		out.PrintCheckSuccess("OpenStack cloud credential")
	} else if _, isValidationProblem := err.(*checker.CheckerOpenstackBadRequest); isValidationProblem {
		return cmderr.CheckFailureError("OpenStack cloud credential")
	}

	return
}
