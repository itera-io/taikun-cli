package complete

import (
	"context"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

// KubernetesVersionCompletionFunc Returns list of Taikun supported Kubernetes versions for a project
func KubernetesVersionCompletionFunc(cmd *cobra.Command, args []string, toComplete string) []string {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Execute a query into the API + graceful exit
	data, _, err := myApiClient.Client.KubernetesAPI.KubernetesGetSupportedList(context.TODO()).Execute()
	if err != nil {
		return []string{}
	}

	// Manipulate the gathered data
	completions := make([]string, 0)

	for i := 0; i < len(data); i++ {
		completions = append(completions, data[i].GetVersion())
	}

	return completions

}
