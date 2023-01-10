package enable

import (
	"errors"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/autoscaling"
	"github.com/itera-io/taikungoclient/client/flavors"
	"github.com/itera-io/taikungoclient/client/servers"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type EnableOptions struct {
	ProjectID            int32
	AutoscalingGroupName string
	Flavor               string
	DiskSize             float64
	MaxSize              int32
	MinSize              int32
}

func NewCmdEnable() *cobra.Command {
	var opts EnableOptions

	cmd := cobra.Command{
		Use:   "enable <project-id>",
		Short: "Enable autoscaling for a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return enableRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.AutoscalingGroupName, "autoscaler-name", "n", "", "The autoscaler's name (required)")
	cmdutils.MarkFlagRequired(&cmd, "autoscaler-name")

	cmd.Flags().StringVarP(&opts.Flavor, "autoscaler-flavor", "f", "", "The autoscaler's flavor (required)")
	cmdutils.MarkFlagRequired(&cmd, "autoscaler-flavor")
	cmdutils.SetFlagCompletionFunc(&cmd, "autoscaler-flavor", flavorCompletionFunc)

	cmd.Flags().Int32Var(&opts.MaxSize, "max-size", 1, "The autoscaler's maximum size")
	cmd.Flags().Int32Var(&opts.MaxSize, "min-size", 1, "The autoscaler's minimum size")
	cmd.Flags().Float64Var(&opts.DiskSize, "disk-size", 30, "The autoscaler's disk size")

	return &cmd
}

func enableRun(opts *EnableOptions) (err error) {
	_, err = isAutoscalingEnabled(opts.ProjectID)
	if err != nil {
		return
	}

	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	body := models.EnableAutoscalingCommand{
		ID:                   opts.ProjectID,
		AutoscalingGroupName: opts.AutoscalingGroupName,
		Flavor:               opts.Flavor,
		DiskSize:             float64(types.GiBToB(int(opts.DiskSize))),
		MaxSize:              opts.MaxSize,
		MinSize:              opts.MinSize,
	}

	params := autoscaling.NewAutoscalingEnableAutoscalingParams().WithV(taikungoclient.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.Autoscaling.AutoscalingEnableAutoscaling(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}
	return
}

func isAutoscalingEnabled(projectID int32) (res bool, err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	params := servers.NewServersDetailsParams().WithV(taikungoclient.Version)
	params = params.WithProjectID(projectID)

	response, err := apiClient.Client.Servers.ServersDetails(params, apiClient)
	if err == nil {
		res := response.Payload.Project.IsAutoscalingEnabled
		if res {
			err = errors.New("Project autoscaling already enabled.")
		}
	}
	return
}

func flavorCompletionFunc(cmd *cobra.Command, args []string, toComplete string) []string {
	if len(args) == 0 {
		return []string{}
	}

	projectID, err := types.Atoi32(args[0])
	if err != nil {
		return []string{}
	}

	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return []string{}
	}

	params := flavors.NewFlavorsGetSelectedFlavorsForProjectParams().WithV(taikungoclient.Version)
	params = params.WithProjectID(&projectID)

	completions := make([]string, 0)

	for {
		response, err := apiClient.Client.Flavors.FlavorsGetSelectedFlavorsForProject(params, apiClient)
		if err != nil {
			return []string{}
		}

		for _, flavor := range response.Payload.Data {
			completions = append(completions, flavor.Name)
		}

		count := int32(len(completions))

		if count == response.Payload.TotalCount {
			break
		}

		params = params.WithOffset(&count)
	}

	return completions
}
