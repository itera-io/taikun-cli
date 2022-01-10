package info

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/format"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/servers"
	"github.com/spf13/cobra"
)

type InfoOptions struct {
	ProjectID int32
}

func NewCmdInfo() *cobra.Command {
	var opts InfoOptions

	cmd := cobra.Command{
		Use:   "info <project-id>",
		Short: "Get detailed information on a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.IDArgumentNotANumberError
			}
			return infoRun(&opts)
		},
	}

	return &cmd
}

func infoRun(opts *InfoOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := servers.NewServersDetailsParams().WithV(apiconfig.Version)
	params = params.WithProjectID(opts.ProjectID)

	response, err := apiClient.Client.Servers.ServersDetails(params, apiClient)
	if err == nil {
		format.PrintResultVertical(response.Payload.Project,
			"projectName",
			"projectStatus",
			"cloudType",
			"kubernetesCurrentVersion",
			"accessProfileName",
			"cloudName",
			"kubernetesProfileName",
			"accessIp",
			"projectHealth",
			"isBackupEnabled",
			"isOpaEnabled",
			"isMonitoringEnabled",
			"hasAlertingProfile",
			"isLocked",
		)
	}

	return
}
