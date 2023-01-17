package complete

import (
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/kubernetes"
	"github.com/spf13/cobra"
)

// Returns list of Taikun supported Kubernetes versions for a project
func KubernetesVersionCompletionFunc(cmd *cobra.Command, args []string, toComplete string) []string {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return []string{}
	}

	params := kubernetes.NewKubernetesGetSupportedListParams().WithV(taikungoclient.Version)

	response, err := apiClient.Client.Kubernetes.KubernetesGetSupportedList(params, apiClient)
	if err != nil {
		return []string{}
	}

	completions := make([]string, 0)

	for i := 0; i < len(response.Payload); i++ {
		completions = append(completions, response.Payload[i].Version)
	}

	return completions
}
