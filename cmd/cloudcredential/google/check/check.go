package check

import (
	"os"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/checker"
	"github.com/spf13/cobra"
)

type CheckOptions struct {
	ConfigFilePath string
}

func NewCmdCheck() *cobra.Command {
	var opts CheckOptions

	cmd := cobra.Command{
		Use:   "check <google-credential-filename>",
		Short: "Check the validity of a Google Cloud Platform credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ConfigFilePath = args[0]
			return checkRun(&opts)
		},
	}

	return &cmd
}

func checkRun(opts *CheckOptions) (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return err
	}

	configFile, err := os.Open(opts.ConfigFilePath)
	if err != nil {
		return err
	}

	params := checker.NewCheckerGoogleParams().WithV(taikungoclient.Version)
	params = params.WithConfig(configFile)

	_, err = apiClient.Client.Checker.CheckerGoogle(params, apiClient)
	if err == nil {
		out.PrintCheckSuccess("Google Cloud Platform credential")
	} else if _, isValidationProblem := err.(*checker.CheckerGoogleBadRequest); isValidationProblem {
		return cmderr.ErrCheckFailure("Google Cloud Platform credential")
	}

	return
}
