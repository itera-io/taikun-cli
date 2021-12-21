package create

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/format"

	"github.com/itera-io/taikungoclient/client/access_profiles"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type CreateOptions struct {
	Name           string
	HttpProxy      string
	OrganizationID int32
	DNSServers     []string
	NTPServers     []string
	IDOnly         bool
}

func NewCmdCreate() *cobra.Command {
	var opts CreateOptions

	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create an access profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			return createRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.HttpProxy, "http-proxy", "p", "", "Http Proxy URL")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")
	cmd.Flags().StringSliceVar(&opts.DNSServers, "dns-servers", []string{}, "DNS Servers")
	cmd.Flags().StringSliceVar(&opts.NTPServers, "ntp-servers", []string{}, "NTP Servers")

	cmdutils.AddIdOnlyFlag(cmd, &opts.IDOnly)

	return cmd
}

func printResult(resource interface{}) {
	if config.OutputFormat == config.OutputFormatJson {
		format.PrettyPrintJson(resource)
	} else if config.OutputFormat == config.OutputFormatTable {
		format.PrettyPrintApiResponseTable(resource,
			"id",
			"name",
			"organizationName",
			"isLocked",
		)
	}
}

func createRun(opts *CreateOptions) (err error) {
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

	params := access_profiles.NewAccessProfilesCreateParams().WithV(apiconfig.Version).WithBody(body)
	response, err := apiClient.Client.AccessProfiles.AccessProfilesCreate(params, apiClient)
	if err == nil {
		if opts.IDOnly {
			format.PrintResourceID(response.Payload)
		} else {
			printResult(response.Payload)
		}
	}

	return
}
