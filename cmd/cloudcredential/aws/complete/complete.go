package complete

import (
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/aws"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

func MakeAwsRegionCompletionFunc(accessKeyID *string, secretAccessKey *string) cmdutils.CompletionCoreFunc {
	return func(cmd *cobra.Command, args []string, toComplete string) (completions []string) {
		completions = make([]string, 0)

		if *accessKeyID == "" || *secretAccessKey == "" {
			return
		}

		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		body := models.RegionListCommand{
			AwsAccessKeyID:     *accessKeyID,
			AwsSecretAccessKey: *secretAccessKey,
		}

		params := aws.NewAwsRegionListParams().WithV(taikungoclient.Version)
		params = params.WithBody(&body)

		result, err := apiClient.Client.Aws.AwsRegionList(params, apiClient)
		if err != nil {
			return
		}

		for _, region := range result.Payload {
			completions = append(completions, region.Region)
		}

		return
	}
}
