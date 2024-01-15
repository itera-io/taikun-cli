package complete

import (
	"context"
	"fmt"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

func MakeAwsRegionCompletionFunc(accessKeyID *string, secretAccessKey *string) cmdutils.CompletionCoreFunc {
	return func(cmd *cobra.Command, args []string, toComplete string) (completions []string) {
		completions = make([]string, 0)

		if *accessKeyID == "" || *secretAccessKey == "" {
			return
		}

		myApiClient := tk.NewClient()

		body := taikuncore.RegionListCommand{
			AwsAccessKeyId:     *taikuncore.NewNullableString(accessKeyID),
			AwsSecretAccessKey: *taikuncore.NewNullableString(secretAccessKey),
		}

		data, response, err := myApiClient.Client.AWSCloudCredentialAPI.AwsRegionlist(context.TODO()).RegionListCommand(body).Execute()
		if err != nil {
			//err = tk.CreateError(response, err)
			fmt.Println(fmt.Errorf(tk.CreateError(response, err).Error())) // This function does not return an error... so just call it. #FIXME
			return
		}

		for _, region := range data {
			completions = append(completions, region.GetRegion())
		}

		return
	}

}
