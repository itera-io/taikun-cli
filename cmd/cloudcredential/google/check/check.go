package check

import (
	"context"
	"fmt"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
	"os"
	"strings"
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
	myApiClient := tk.NewClient()
	configFile, err := os.Open(opts.ConfigFilePath)
	if err != nil {
		return err
	}
	_, response, err := myApiClient.Client.CheckerAPI.CheckerGoogle(context.TODO()).Config(configFile).Execute()
	if err == nil {
		out.PrintCheckSuccess("Google Cloud Platform credential")
	}

	// Did it fail because the request failed (e.g. cannot connect to Taikun) or because the credentials are not valid?
	if err != nil {
		myError := tk.CreateError(response, err)
		myStringError := fmt.Sprint(myError)
		if strings.Contains(myStringError, "Failed to validate") {
			err = cmderr.ErrCheckFailure("Google Cloud Platform credential") // Taikun responded that credentials are not valid.
		} else {
			err = tk.CreateError(response, err) // Something else happened
		}

		return
	}

	return

	/*
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
	*/
}
