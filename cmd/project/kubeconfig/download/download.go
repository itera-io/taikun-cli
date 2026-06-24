package download

import (
	"fmt"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
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
			return downloadRun(cmd, &opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.KubeconfigID, "kubeconfig-id", "k", 0, "Kubeconfig ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "kubeconfig-id")

	cmd.Flags().StringVarP(&opts.OutputFile, "output", "o", "", "Output filename")

	return &cmd
}

func downloadRun(cmd *cobra.Command, opts *DownloadOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()

	if opts.OutputFile == "" {
		kubeconfigName, err := getKubeconfigName(cmd, opts)
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
	data, response, err := myApiClient.Client.KubeConfigAPI.KubeconfigDownload(ctx).DownloadKubeConfigCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	content := []byte(data)

	return os.WriteFile(opts.OutputFile, content, 0644)

}

func getKubeconfigName(cmd *cobra.Command, opts *DownloadOptions) (name string, err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.KubeConfigAPI.KubeconfigList(ctx).ProjectId(opts.ProjectID).Id(opts.KubeconfigID).Execute()
	//data, response, err := myApiClient.Client.KubeConfigAPI.KubeconfigList(ctx).Id(kubeconfigID).Execute()
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

}
