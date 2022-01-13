package download

import (
	"fmt"
	"os"

	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/kube_config"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type DownloadOptions struct {
	KubeconfigID int32
	ProjectID    int32
	OutputFile   string
}

func NewCmdDownload() *cobra.Command {
	var opts DownloadOptions

	cmd := cobra.Command{
		Use:   "download <kubeconfig-id>",
		Short: "Download a kubeconfig",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.KubeconfigID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return downloadRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.ProjectID, "project-id", "p", 0, "Project ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "project-id")

	cmd.Flags().StringVarP(&opts.OutputFile, "output", "o", "", "Output filename")

	return &cmd
}

func downloadRun(opts *DownloadOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	if opts.OutputFile == "" {
		kubeconfigName, err := getKubeconfigName(opts.KubeconfigID)
		if err != nil {
			return err
		}
		opts.OutputFile = fmt.Sprintf(
			"taikun-%d-%s.yaml",
			opts.ProjectID,
			kubeconfigName,
		)
	}

	body := models.DownloadKubeConfigCommand{
		ID:        opts.KubeconfigID,
		ProjectID: opts.ProjectID,
	}

	params := kube_config.NewKubeConfigDownloadParams().WithV(api.Version)
	params = params.WithBody(&body)

	response, err := apiClient.Client.KubeConfig.KubeConfigDownload(params, apiClient)
	if err == nil {
		content := []byte(response.Payload.(string))
		err = os.WriteFile(opts.OutputFile, content, 0644)
	}

	return
}

func getKubeconfigName(id int32) (name string, err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := kube_config.NewKubeConfigListParams().WithV(api.Version)
	params = params.WithID(&id)

	response, err := apiClient.Client.KubeConfig.KubeConfigList(params, apiClient)
	if err != nil {
		return
	}
	if len(response.Payload.Data) != 1 {
		return "", cmderr.ResourceNotFoundError("Kubeconfig", id)
	}

	name = response.Payload.Data[0].ServiceAccountName

	return
}