package add

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

var addFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"NAME", "name",
		),
		field.NewVisible(
			"ORG", "organizationName",
		),
		field.NewHidden(
			"ORG-ID", "organizationId",
		),
		field.NewVisible(
			"TAIKUN-LB", "taikunLBEnabled",
		),
		field.NewVisible(
			"OCTAVIA", "octaviaEnabled",
		),
		field.NewVisible(
			"BASTION-PROXY", "exposeNodePortOnBastion",
		),
		field.NewVisible(
			"CNI", "cni",
		),
		field.NewVisible(
			"SCHEDULE-ON-MASTER", "allowSchedulingOnMaster",
		),
		field.NewVisibleWithToStringFunc(
			"LOCK", "isLocked", out.FormatLockStatus,
		),
		field.NewHidden(
			"CREATED-BY", "createdBy",
		),
	},
)

type AddOptions struct {
	AllowSchedulingOnMaster  bool
	ExposeNodePortOnBastion  bool
	Name                     string
	OctaviaEnabled           bool
	OrganizationID           int32
	TaikunLBEnabled          bool
	DisableUniqueClusterName bool
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <name>",
		Short: "Add a kubernetes profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			return addRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")
	cmd.Flags().BoolVar(&opts.AllowSchedulingOnMaster, "allow-master-scheduling", false, "Allow scheduling on master nodes")
	cmd.Flags().BoolVar(&opts.ExposeNodePortOnBastion, "expose-node-port-on-bastion", false, "Expose Node Port on Bastion")
	cmd.Flags().BoolVar(&opts.OctaviaEnabled, "enable-octavia", false, "Enable Octavia Load Balancer")
	cmd.Flags().BoolVar(&opts.TaikunLBEnabled, "enable-taikun-lb", false, "Enable Taikun Load Balancer")
	cmd.Flags().BoolVar(&opts.DisableUniqueClusterName, "disable-unique-cluster-name", false, "Disable unique cluster name, the cluster name will be cluster.local")

	cmdutils.AddOutputOnlyIDFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, addFields)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	useUniqueName := !opts.DisableUniqueClusterName
	body := taikuncore.CreateKubernetesProfileCommand{
		Name:                    *taikuncore.NewNullableString(&opts.Name),
		OctaviaEnabled:          &opts.OctaviaEnabled,
		ExposeNodePortOnBastion: &opts.ExposeNodePortOnBastion,
		OrganizationId:          *taikuncore.NewNullableInt32(&opts.OrganizationID),
		TaikunLBEnabled:         &opts.TaikunLBEnabled,
		AllowSchedulingOnMaster: &opts.AllowSchedulingOnMaster,
		UniqueClusterName:       &useUniqueName,
	}

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.KubernetesProfilesAPI.KubernetesprofilesCreate(context.TODO()).CreateKubernetesProfileCommand(body).Execute()
	if err != nil {
		err = tk.CreateError(response, err)
		return
	}

	// Manipulate the gathered data
	return out.PrintResult(data, addFields)

}
