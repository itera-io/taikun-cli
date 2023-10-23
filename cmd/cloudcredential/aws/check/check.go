package check

import (
	"context"
	"fmt"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/aws/complete"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"strings"
)

type CheckOptions struct {
	AWSSecretAccessKey string
	AWSAccessKeyID     string
	AWSRegion          string
}

func NewCmdCheck() *cobra.Command {
	var opts CheckOptions

	cmd := &cobra.Command{
		Use:   "check",
		Short: "Check the validity of an AWS cloud credential",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return checkRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.AWSSecretAccessKey, "secret-access-key", "s", "", "AWS Secret Access Key (required)")
	cmdutils.MarkFlagRequired(cmd, "secret-access-key")

	cmd.Flags().StringVarP(&opts.AWSAccessKeyID, "access-key-id", "a", "", "AWS Access Key ID (required)")
	cmdutils.MarkFlagRequired(cmd, "access-key-id")

	cmd.Flags().StringVarP(&opts.AWSRegion, "region", "r", "", "AWS Region (required)")
	cmdutils.MarkFlagRequired(cmd, "region")
	cmdutils.SetFlagCompletionFunc(cmd, "region", complete.MakeAwsRegionCompletionFunc(&opts.AWSAccessKeyID, &opts.AWSSecretAccessKey))

	return cmd
}

func checkRun(opts *CheckOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.CheckAwsCommand{
		AwsAccessKeyId:     *taikuncore.NewNullableString(&opts.AWSAccessKeyID),
		AwsSecretAccessKey: *taikuncore.NewNullableString(&opts.AWSSecretAccessKey),
		Region:             *taikuncore.NewNullableString(&opts.AWSRegion),
	}

	// Execute a query into the API + graceful exit
	myRequest := myApiClient.Client.CheckerAPI.CheckerAws(context.TODO()).CheckAwsCommand(body)
	response, err := myRequest.Execute()

	if err == nil {
		out.PrintCheckSuccess("AWS cloud credential")
	}

	// Did it fail because the request failed (e.g. cannot connect to Taikun) or because the credentials are not valid?
	if err != nil {
		myError := tk.CreateError(response, err)
		myStringError := fmt.Sprint(myError)
		if strings.Contains(myStringError, "Failed to validate") {
			err = cmderr.ErrCheckFailure("AWS cloud credential") // Taikun responded that credentials are not valid.
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

		body := models.CheckAwsCommand{
			AwsSecretAccessKey: opts.AWSSecretAccessKey,
			AwsAccessKeyID:     opts.AWSAccessKeyID,
			Region:             opts.AWSRegion,
		}

		params := checker.NewCheckerAwsParams().WithV(taikungoclient.Version).WithBody(&body)

		_, err = apiClient.Client.Checker.CheckerAws(params, apiClient)
		if err == nil {
			out.PrintCheckSuccess("AWS cloud credential")
		} else if _, isValidationProblem := err.(*checker.CheckerAwsBadRequest); isValidationProblem {
			return cmderr.ErrCheckFailure("AWS cloud credential")
		}

		return
	*/
}
