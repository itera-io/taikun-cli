package add

import (
	"errors"
	"fmt"
	"strings"

	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/servers"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type AddOptions struct {
	DiskSize             int
	Flavor               string
	KubernetesNodeLabels []string
	Name                 string
	ProjectID            int32
	Role                 string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <name>",
		Short: "Add a server",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			if err := cmdutils.CheckFlagValue("role", opts.Role, types.ServerRoles); err != nil {
				return err
			}
			return addRun(&opts)
		},
	}

	cmd.Flags().IntVarP(&opts.DiskSize, "disk-size", "d", 30, "Disk size in GB")

	cmd.Flags().StringVarP(&opts.Flavor, "flavor", "f", "", "Flavor (required)")
	cmdutils.MarkFlagRequired(&cmd, "flavor")

	cmd.Flags().StringSliceVarP(&opts.KubernetesNodeLabels, "kubernetes-node-labels", "k", []string{}, "Kubernetes node labels (format: \"key=value,key2=value2,...\")")

	cmd.Flags().StringVarP(&opts.Role, "role", "r", "", "Role (required)")
	cmdutils.MarkFlagRequired(&cmd, "role")
	cmdutils.RegisterStaticFlagCompletion(&cmd, "role", types.ServerRoles.Keys()...)

	cmd.Flags().Int32VarP(&opts.ProjectID, "project-id", "p", 0, "Project ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "project-id")

	cmdutils.AddOutputOnlyIDFlag(&cmd)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.ServerForCreateDto{
		DiskSize:  types.GiBToB(opts.DiskSize),
		Flavor:    opts.Flavor,
		Name:      opts.Name,
		ProjectID: opts.ProjectID,
		Role:      types.GetServerRole(opts.Role),
	}

	if len(opts.KubernetesNodeLabels) != 0 {
		body.KubernetesNodeLabels, err = parseKubernetesNodeLabelsFlag(opts.KubernetesNodeLabels)
		if err != nil {
			return
		}
	}

	params := servers.NewServersCreateParams().WithV(apiconfig.Version)
	params = params.WithBody(&body)

	response, err := apiClient.Client.Servers.ServersCreate(params, apiClient)
	if err == nil {
		out.PrintResult(response.Payload,
			"id",
			"name",
			"cpu",
			"ram",
			"diskSize",
			"role",
			"status",
		)
	}

	return
}

func parseKubernetesNodeLabelsFlag(labelsData []string) ([]*models.KubernetesNodeLabelsDto, error) {
	labels := make([]*models.KubernetesNodeLabelsDto, len(labelsData))
	for i, labelData := range labelsData {
		if len(labelData) == 0 {
			return nil, errors.New("Invalid empty kubernetes node label")
		}
		tokens := strings.Split(labelData, "=")
		if len(tokens) != 2 {
			return nil, fmt.Errorf("Invalid kubernetes node label format: %s", labelData)
		}
		labels[i] = &models.KubernetesNodeLabelsDto{
			Key:   tokens[0],
			Value: tokens[1],
		}
	}
	return labels, nil
}
