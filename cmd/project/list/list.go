package list

import (
	"context"
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
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
			"STATUS", "status",
		),
		field.NewVisibleWithToStringFunc(
			"HEALTH", "health", out.FormatProjectHealth,
		),
		field.NewHiddenWithToStringFunc(
			"CREATED-AT", "createdAt", out.FormatDateTimeString,
		),
		field.NewVisibleWithToStringFunc(
			"CLOUD", "cloudType", out.FormatCloudType,
		),
		field.NewVisible("VIRTUAL", "isVirtualCluster"),
		field.NewVisible(
			"K8S", "isKubernetes",
		),
		field.NewVisible(
			"QUOTA-ID", "quotaId",
		),
		field.NewHidden(
			"CLOUD-CREDENTIAL", "cloudCredentialName",
		),
		field.NewHidden(
			"HAS-KUBECONFIG", "hasKubeConfigFile",
		),
		field.NewHidden(
			"K8S-VERSION", "kubernetesCurrentVersion",
		),
		field.NewHidden(
			"KUBESPRAY-VERSION", "kubesprayCurrentVersion",
		),
		field.NewHidden(
			"SERVERS", "totalServersCount",
		),
		field.NewVisibleWithToStringFunc(
			"EXPIRES", "expiredAt", out.FormatDateTimeString,
		),
		field.NewHidden(
			"DELETE-ON-EXPIRATION", "deleteOnExpiration",
		),
		field.NewVisible(
			"LOCK", "isLocked",
		),
		field.NewHiddenWithToStringFunc(
			"LAST-MODIFIED", "lastModified", out.FormatDateTimeString,
		),
		field.NewHidden(
			"LAST-MODIFIED-BY", "lastModifiedBy",
		),
		field.NewVisible(
			"AUTOSCALER", "isAutoscalingEnabled",
		),
	},
)

type ListOptions struct {
	OrganizationID int32
	Limit          int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list",
		Short: "List projects",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role). Default 0.")

	cmdutils.AddSortByAndReverseFlags(&cmd, "projects", listFields)
	cmdutils.AddLimitFlag(&cmd, &opts.Limit)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.ProjectsAPI.ProjectsList(context.TODO())
	if opts.OrganizationID != 0 {
		myRequest = myRequest.OrganizationId(opts.OrganizationID)
	}
	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}

	var projects = make([]taikuncore.ProjectListDetailDto, 0)

	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			return tk.CreateError(response, err)
		}
		projects = append(projects, data.GetData()...)

		projectsCount := int32(len(projects))
		if opts.Limit != 0 && projectsCount >= opts.Limit {
			break
		}

		if projectsCount == data.GetTotalCount() {
			break
		}

		myRequest = myRequest.Offset(projectsCount)
	}

	if opts.Limit != 0 && int32(len(projects)) > opts.Limit {
		projects = projects[:opts.Limit]
	}

	return out.PrintResults(projects, listFields)

}
