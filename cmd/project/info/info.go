package info

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

var infoFields = fields.New(
	[]*field.Field{
		field.NewHidden(
			"ID", "id",
		),
		field.NewVisible(
			"NAME", "name",
		),
		field.NewVisible(
			"ORG", "organizationName",
		),
		field.NewVisible(
			"ORG-ID", "organizationId",
		),
		field.NewVisibleWithToStringFunc(
			"HEALTH", "health", out.FormatProjectHealth,
		),
		field.NewVisible(
			"STATUS", "status",
		),
		field.NewVisible(
			"ACCESS-PROFILE", "accessProfileName",
		),
		field.NewVisible(
			"ACCESS-PROFILE-ID", "accessProfileId",
		),
		field.NewVisibleWithToStringFunc(
			"CLOUD", "cloudType", out.FormatCloudType,
		),
		field.NewVisible(
			"CLOUD-CREDENTIAL", "cloudName",
		),
		field.NewVisible(
			"CLOUD-CREDENTIAL-ID", "cloudId",
		),
		field.NewVisible(
			"ALERTING-PROFILE", "alertingProfileName",
		),
		field.NewVisible(
			"ALERTING-PROFILE-ID", "alertingProfileId",
		),
		field.NewVisible(
			"AUTO-UPGRADES", "isAutoUpgrade",
		),
		// No longer in the API
		//field.NewVisible(
		//	"UPGRADABLE", "hasNextVersion",
		//),
		field.NewVisible(
			"HAS-FLAVORS", "hasSelectedFlavors",
		),
		field.NewVisible(
			"BACKUPS", "isBackupEnabled",
		),
		field.NewVisible(
			"MAINTENANCE", "isMaintenanceModeEnabled",
		),
		field.NewVisible(
			"MONITORING", "isMonitoringEnabled",
		),
		field.NewVisible(
			"POLICY-PROFILE", "opaProfileName",
		),
		field.NewVisible(
			"POLICY-PROFILE-ID", "opaProfileId",
		),
		field.NewVisible(
			"HAS-KUBECONFIG", "hasKubeConfigFile",
		),
		field.NewVisible(
			"K8S-VERSION", "kubernetesVersion",
		),
		field.NewHidden(
			"KUBESPRAY-VERSION", "kubeCurrentVersion",
		),
		field.NewVisible(
			"K8S-PROFILE-ID", "kubernetesProfileId",
		),
		field.NewVisible(
			"K8S-PROFILE", "kubernetesProfileName",
		),
		field.NewVisible(
			"QUOTA-ID", "quotaId",
		),
		// Removed from the API
		//field.NewVisible(
		//	"REVISIONS", "projectRevision",
		//),
		//field.NewVisible(
		//	"SERVERS", "totalCount",
		//),
		field.NewVisible(
			"BASTIONS", "bastion",
		),
		field.NewVisible(
			"KUBEMASTERS", "masterReady",
		),
		field.NewVisible(
			"KUBEWORKERS", "worker",
		),
		field.NewVisible(
			"TOTAL-CPU", "usedCpu",
		),
		field.NewVisible(
			"TOTAL-DISK", "usedDiskSize",
		),
		field.NewVisible(
			"TOTAL-RAM", "usedRam",
		),
		field.NewVisibleWithToStringFunc(
			"LOCK", "isLocked", out.FormatLockStatus,
		),
		field.NewVisible(
			"WASM", "wasmEnabled",
		),
		field.NewVisible(
			"SPOT", "spotEnabled",
		),
		field.NewHidden(
			"SPOT-FULL", "allowFullSpotKubernetes",
		),
		field.NewHidden(
			"SPOT-WORKERS", "allowSpotWorkers",
		),
		field.NewHidden(
			"SPOT-VMS", "allowSpotVMs",
		),
		field.NewVisible(
			"AUTOSCALING", "isAutoscalingEnabled",
		),
		field.NewHidden(
			"AUTOSCALING-GROUP", "autoscalingGroupName",
		),
		field.NewHidden(
			"AUTOSCALING-MIN", "minSize",
		),
		field.NewHidden(
			"AUTOSCALING-MAX", "maxSize",
		),
		field.NewHidden(
			"AUTOSCALING-FLAVOR", "flavor",
		),
		field.NewHidden(
			"AUTOSCALING-SPOT", "isAutoscalingSpotEnabled",
		),
	},
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
				return cmderr.ErrIDArgumentNotANumber
			}
			return infoRun(&opts)
		},
	}

	cmdutils.AddColumnsFlag(&cmd, infoFields)

	return &cmd
}

func infoRun(opts *InfoOptions) (err error) {
	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.ServersAPI.ServersDetails(context.TODO(), opts.ProjectID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	myProject := data.GetProject()

	// Show if any of the spots is enabled for clarity
	myProject.SetSpotEnabled(myProject.GetAllowSpotVMs() || myProject.GetAllowSpotWorkers() || myProject.GetAllowFullSpotKubernetes())

	return out.PrintResult(myProject, infoFields)

}
