package create

import (
	"taikun-cli/api"
	"taikun-cli/cmd/cmdutils"

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
}

func NewCmdCreate() *cobra.Command {
	var opts CreateOptions

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create an access profile",
		RunE: func(cmd *cobra.Command, args []string) error {
			return createRun(&opts)
		},
	}

	cmd.Flags().StringVar(&opts.Name, "name", "", "Name (required)")
	cmdutils.MarkFlagRequired(cmd, "name")

	cmd.Flags().StringVar(&opts.HttpProxy, "http-proxy", "", "HttpProxy")
	cmd.Flags().Int32Var(&opts.OrganizationID, "organization-id", 0, "Organization ID")
	cmd.Flags().StringSliceVar(&opts.DNSServers, "dns-servers", []string{}, "DNS Servers")
	cmd.Flags().StringSliceVar(&opts.NTPServers, "ntp-servers", []string{}, "NTP Servers")

	return cmd
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

	params := access_profiles.NewAccessProfilesCreateParams().WithV(cmdutils.ApiVersion).WithBody(body)
	response, err := apiClient.Client.AccessProfiles.AccessProfilesCreate(params, apiClient)
	if err == nil {
		cmdutils.PrettyPrint(response.Payload)
	}

	return
}
