package cmdutils

import (
	"context"
	tk "github.com/Smidra/taikungoclient"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/spf13/cobra"
)

func FlavorCompletionFunc(cmd *cobra.Command, args []string, toComplete string) []string {
	if len(args) == 0 {
		return []string{}
	}

	projectID, err := types.Atoi32(args[0])
	if err != nil {
		return []string{}
	}
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.FlavorsAPI.FlavorsSelectedFlavorsForProject(context.TODO()).ProjectId(projectID)
	completions := make([]string, 0)

	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			err = tk.CreateError(response, err)
			return []string{}
		}

		for _, flavor := range data.Data {
			completions = append(completions, flavor.GetName())
		}

		count := int32(len(completions))

		if count == data.GetTotalCount() {
			break
		}

		myRequest = myRequest.Offset(count)
	}

	return completions
}
