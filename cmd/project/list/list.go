package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikungoclient/client/projects"
	"github.com/itera-io/taikungoclient/models"
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
		field.NewVisible(
			"LOCK", "isLocked",
		),
		field.NewHiddenWithToStringFunc(
			"LAST-MODIFIED", "lastModified", out.FormatDateTimeString,
		),
		field.NewHidden(
			"LAST-MODIFIED-BY", "lastModifiedBy",
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

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddSortByAndReverseFlags(&cmd, "projects", listFields)
	cmdutils.AddLimitFlag(&cmd, &opts.Limit)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := projects.NewProjectsListParams().WithV(api.Version)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}
	if config.SortBy != "" {
		params = params.WithSortBy(&config.SortBy).WithSortDirection(api.GetSortDirection())
	}

	var projects = make([]*models.ProjectListForUIDto, 0)
	for {
		response, err := apiClient.Client.Projects.ProjectsList(params, apiClient)
		if err != nil {
			return err
		}
		projects = append(projects, response.Payload.Data...)
		projectsCount := int32(len(projects))
		if opts.Limit != 0 && projectsCount >= opts.Limit {
			break
		}
		if projectsCount == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&projectsCount)
	}

	if opts.Limit != 0 && int32(len(projects)) > opts.Limit {
		projects = projects[:opts.Limit]
	}

	return out.PrintResults(projects, listFields)
}
