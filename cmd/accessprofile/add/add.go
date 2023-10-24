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
		field.NewVisible(
			"ORG-ID", "organizationId",
		),
		field.NewVisible(
			"HTTP-PROXY", "httpProxy",
		),
		field.NewHiddenWithToStringFunc(
			"LOCK", "isLocked", out.FormatLockStatus,
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

func addRun(opts *AddOptions) error {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	DNSServers := make([]taikuncore.DnsServerCreateDto, len(opts.DNSServers))
	for i, rawDNSServer := range opts.DNSServers {
		DNSServers[i] = taikuncore.DnsServerCreateDto{
			Address: *taikuncore.NewNullableString(&rawDNSServer),
		}
	}
	NTPServers := make([]taikuncore.NtpServerCreateDto, len(opts.NTPServers))
	for i, rawNTPServer := range opts.NTPServers {
		NTPServers[i] = taikuncore.NtpServerCreateDto{
			Address: *taikuncore.NewNullableString(&rawNTPServer),
		}
	}
	body := taikuncore.CreateAccessProfileCommand{
		Name:           *taikuncore.NewNullableString(&opts.Name),
		OrganizationId: *taikuncore.NewNullableInt32(&opts.OrganizationID),
		DnsServers:     DNSServers,
		NtpServers:     NTPServers,
	}
	if opts.HttpProxy != "" {
		body.SetHttpProxy(opts.HttpProxy)
	}

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.AccessProfilesAPI.AccessprofilesCreate(context.TODO()).CreateAccessProfileCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	// Manipulate the gathered data
	return out.PrintResult(data, addFields)

}
