package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/servers"
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
			"IP", "ipAddress",
		),
		field.NewHidden(
			"CLOUD", "cloudType",
		),
		field.NewVisible(
			"FLAVOR", "",
			// JSON property name is set in the listRun function
			// as it depends on the server's cloud type
		),
		field.NewVisible(
			"CPU", "cpu",
		),
		field.NewVisibleWithToStringFunc(
			"RAM", "ram", out.FormatBToGiB,
		),
		field.NewVisibleWithToStringFunc(
			"DISK", "diskSize", out.FormatBToGiB,
		),
		field.NewVisible(
			"ROLE", "role",
		),
		field.NewVisible(
			"STATUS", "status",
		),
		field.NewHidden(
			"PROJECT", "projectName",
		),
		field.NewHidden(
			"PROJECT-ID", "projectId",
		),
		field.NewHidden(
			"ORG", "organizationName",
		),
		field.NewHidden(
			"ORG-ID", "organizationId",
		),
		field.NewVisibleWithToStringFunc(
			"CREATED-AT", "createdAt", out.FormatDateTimeString,
		),
		field.NewHidden(
			"CREATED-BY", "createdBy",
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
	ProjectID int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list <project-id>",
		Short: "List a project's Kubernetes servers",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmdutils.AddSortByAndReverseFlags(&cmd, "projects-k8s", listFields)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	projectServers, err := ListServers(opts)
	if err == nil {
		flavorJsonPropertyName, err := getFlavorField(projectServers)
		if err != nil {
			return err
		}

		if err := listFields.SetFieldJsonPropertyName("FLAVOR", flavorJsonPropertyName); err != nil {
			return err
		}

		return out.PrintResults(projectServers, listFields)
	}

	return
}

func ListServers(opts *ListOptions) (projectServers []*models.ServerListDto, err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := servers.NewServersDetailsParams().WithV(api.Version)
	params = params.WithProjectID(opts.ProjectID)

	if config.SortBy != "" {
		params = params.WithSortBy(&config.SortBy)
		params = params.WithSortDirection(api.GetSortDirection())
	}

	response, err := apiClient.Client.Servers.ServersDetails(params, apiClient)
	if err == nil {
		projectServers = response.Payload.Data
	}

	return
}

func getFlavorField(servers []*models.ServerListDto) (string, error) {
	if len(servers) == 0 {
		return "flavor", nil
	}

	if servers[0].AwsInstanceType != "" {
		return "awsInstanceType", nil
	}

	if servers[0].AzureVMSize != "" {
		return "azureVmSize", nil
	}

	if servers[0].OpenstackFlavor != "" {
		return "openstackFlavor", nil
	}

	return "", cmderr.ServerHasNoFlavorError
}
