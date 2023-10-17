package download

import (
	"context"
	"fmt"
	tk "github.com/Smidra/taikungoclient"
	taikuncore "github.com/Smidra/taikungoclient/client"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/spf13/cobra"
	"os"
)

type DownloadOptions struct {
	KubeconfigID int32
	ProjectID    int32
	OutputFile   string
}

func NewCmdDownload() *cobra.Command {
	var opts DownloadOptions

	cmd := cobra.Command{
		Use:   "download <project-id>",
		Short: "Download a kubeconfig",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return downloadRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.KubeconfigID, "kubeconfig-id", "k", 0, "Kubeconfig ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "kubeconfig-id")

	cmd.Flags().StringVarP(&opts.OutputFile, "output", "o", "", "Output filename")

	return &cmd
}

func downloadRun(opts *DownloadOptions) (err error) {
	myApiClient := tk.NewClient()

	if opts.OutputFile == "" {
		kubeconfigName, err := getKubeconfigName(opts)
		if err != nil {
			return err
		}

		opts.OutputFile = fmt.Sprintf(
			"taikun-%d-%s.yaml",
			opts.ProjectID,
			kubeconfigName,
		)
	}
	body := taikuncore.DownloadKubeConfigCommand{
		Id:        &opts.KubeconfigID,
		ProjectId: &opts.ProjectID,
	}
	data, response, err := myApiClient.Client.KubeConfigAPI.KubeconfigDownload(context.TODO()).DownloadKubeConfigCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	content := []byte(data)

	return os.WriteFile(opts.OutputFile, content, 0644)

	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return err
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

		params := kube_config.NewKubeConfigDownloadParams().WithV(taikungoclient.Version)
		params = params.WithBody(&body)

		response, err := apiClient.Client.KubeConfig.KubeConfigDownload(params, apiClient)
		if err != nil {
			return err
		}

		payload, payloadOk := response.Payload.(string)
		if !payloadOk {
			return cmderr.ProgramError("downloadRun", errors.New("failed to convert payload to string"))
		}

		content := []byte(payload)

		return os.WriteFile(opts.OutputFile, content, 0644)
	*/
}

func getKubeconfigName(opts *DownloadOptions) (name string, err error) {
	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.KubeConfigAPI.KubeconfigList(context.TODO()).ProjectId(opts.ProjectID).Id(opts.KubeconfigID).Execute()
	//data, response, err := myApiClient.Client.KubeConfigAPI.KubeconfigList(context.TODO()).Id(kubeconfigID).Execute()
	if err != nil {
		err = tk.CreateError(response, err)
		return
	}
	if len(data.GetData()) != 1 {
		return "", cmderr.ResourceNotFoundError("Kubeconfig", opts.KubeconfigID)
	}

	name = *data.GetData()[0].DisplayName.Get()
	//name = response.Payload.Data[0].DisplayName

	return

	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		params := kube_config.NewKubeConfigListParams().WithV(taikungoclient.Version)
		params = params.WithID(&kubeconfigID)

		response, err := apiClient.Client.KubeConfig.KubeConfigList(params, apiClient)
		if err != nil {
			return
		}

		if len(response.Payload.Data) != 1 {
			return "", cmderr.ResourceNotFoundError("Kubeconfig", kubeconfigID)
		}

		name = response.Payload.Data[0].DisplayName

		return
	*/
}
