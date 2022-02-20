package check

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/aws/complete"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikungoclient/client/checker"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
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
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.CheckAwsCommand{
		AwsSecretAccessKey: opts.AWSSecretAccessKey,
		AwsAccessKeyID:     opts.AWSAccessKeyID,
		Region:             opts.AWSRegion,
	}

	params := checker.NewCheckerAwsParams().WithV(api.Version).WithBody(&body)

	_, err = apiClient.Client.Checker.CheckerAws(params, apiClient)
	if err == nil {
		out.PrintCheckSuccess("AWS cloud credential")
	} else if _, isValidationProblem := err.(*checker.CheckerAwsBadRequest); isValidationProblem {
		return cmderr.ErrCheckFailure("AWS cloud credential")
	}

	return
}
