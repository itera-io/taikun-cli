package add

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fieldnames"
	"github.com/itera-io/taikun-cli/utils/out/fields"

	"github.com/itera-io/taikungoclient/client/access_profiles"
	"github.com/itera-io/taikungoclient/models"
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
		field.NewVisible(
			"ORG-ID", "organizationId",
		),
		field.NewVisible(
			"HTTP-PROXY", "httpProxy",
		),
		field.NewHiddenWithToStringFunc(
			fieldnames.IsLocked, "isLocked", out.FormatLockStatus,
		),
		field.NewHidden(
			"CREATED-BY", "createdBy",
		),
	},
)

type AddOptions struct {
	Name           string
	HttpProxy      string
	OrganizationID int32
	DNSServers     []string
	NTPServers     []string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := &cobra.Command{
		Use:   "add <name>",
		Short: "Add an access profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.HttpProxy, "http-proxy", "p", "", "Http Proxy URL")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")
	cmd.Flags().StringSliceVarP(&opts.DNSServers, "dns-servers", "d", []string{}, "DNS Servers")
	cmd.Flags().StringSliceVarP(&opts.NTPServers, "ntp-servers", "n", []string{}, "NTP Servers")
	cmdutils.AddOutputOnlyIDFlag(cmd)
	cmdutils.AddColumnsFlag(cmd, addFields)

	return cmd
}

func addRun(opts *AddOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	DNSServers := make([]*models.DNSServerListDto, len(opts.DNSServers))
	for i, rawDNSServer := range opts.DNSServers {
		DNSServers[i] = &models.DNSServerListDto{
			Address: rawDNSServer,
		}
	}
	NTPServers := make([]*models.NtpServerListDto, len(opts.NTPServers))
	for i, rawNTPServer := range opts.NTPServers {
		NTPServers[i] = &models.NtpServerListDto{
			Address: rawNTPServer,
		}
	}

	body := &models.UpsertAccessProfileCommand{
		Name:           opts.Name,
		HTTPProxy:      opts.HttpProxy,
		OrganizationID: opts.OrganizationID,
		DNSServers:     DNSServers,
		NtpServers:     NTPServers,
	}

	params := access_profiles.NewAccessProfilesCreateParams().WithV(api.Version).WithBody(body)
	response, err := apiClient.Client.AccessProfiles.AccessProfilesCreate(params, apiClient)
	if err == nil {
		out.PrintResult(response.Payload, addFields)
	}

	return
}
