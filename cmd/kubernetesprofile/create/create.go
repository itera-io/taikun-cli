package create

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"

	"github.com/itera-io/taikungoclient/client/kubernetes_profiles"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type CreateOptions struct {
	AllowSchedulingOnMaster bool
	ExposeNodePortOnBastion bool
	Name                    string
	OctaviaEnabled          bool
	OrganizationID          int32
	TaikunLBEnabled         bool
}

func NewCmdCreate() *cobra.Command {
	var opts CreateOptions

	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create a kubernetes profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			return createRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")
	cmd.Flags().BoolVar(&opts.AllowSchedulingOnMaster, "allow-master-scheduling", false, "Allow scheduling on master nodes")
	cmd.Flags().BoolVar(&opts.ExposeNodePortOnBastion, "expose-node-port-on-bastion", false, "Expose Node Port on Bastion")
	cmd.Flags().BoolVar(&opts.OctaviaEnabled, "enable-octavia", false, "Enable Octavia Load Balancer")
	cmd.Flags().BoolVar(&opts.TaikunLBEnabled, "enable-taikun-lb", false, "Enable Taikun Load Balancer")

	cmdutils.AddOutputOnlyIDFlag(cmd)

	return cmd
}

func createRun(opts *CreateOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := &models.CreateKubernetesProfileCommand{
		AllowSchedulingOnMaster: opts.AllowSchedulingOnMaster,
		ExposeNodePortOnBastion: opts.ExposeNodePortOnBastion,
		Name:                    opts.Name,
		OctaviaEnabled:          opts.OctaviaEnabled,
		OrganizationID:          opts.OrganizationID,
		TaikunLBEnabled:         opts.TaikunLBEnabled,
	}

	params := kubernetes_profiles.NewKubernetesProfilesCreateParams().WithV(apiconfig.Version).WithBody(body)
	response, err := apiClient.Client.KubernetesProfiles.KubernetesProfilesCreate(params, apiClient)
	if err == nil {
		out.PrintResult(response.Payload,
			"id",
			"name",
			"organizationName",
			"taikunLBEnabled",
			"octaviaEnabled",
			"exposeNodePortOnBastion",
			"cni",
			"allowSchedulingOnMaster",
		)
	}

	return
}
