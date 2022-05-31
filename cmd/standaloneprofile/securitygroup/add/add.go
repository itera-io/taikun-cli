package add

import (
	"errors"
	"fmt"
	"strings"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/security_group"
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
			"REMOTE-IP-PREFIX", "remoteIpPrefix",
		),
		field.NewVisibleWithToStringFunc(
			"PROTOCOL", "protocol", out.FormatStringUpper,
		),
		field.NewVisible(
			"MIN-PORT", "portMinRange",
		),
		field.NewVisible(
			"MAX-PORT", "portMaxRange",
		),
	},
)

type AddOptions struct {
	MaxPort             int32
	MinPort             int32
	Name                string
	Protocol            string
	RemoteIpPrefix      string
	StandAloneProfileID int32
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <standalone-profile-id>",
		Short: "Add a security group to a standalone profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.StandAloneProfileID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			if err := cmdutils.CheckFlagValue("protocol", opts.Protocol, types.SecurityGroupProtocols); err != nil {
				return err
			}
			if strings.ToLower(opts.Protocol) == "icmp" {
				if opts.MinPort != -1 || opts.MaxPort != -1 {
					return errors.New("cannot set port range with ICMP protocol")
				}
			} else {
				if opts.MinPort == -1 || opts.MaxPort == -1 {
					return fmt.Errorf("must set --min-port and --max-port with %s protocol", opts.Protocol)
				}
			}
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Name, "name", "n", "", "Name (required)")
	cmdutils.MarkFlagRequired(&cmd, "name")

	cmd.Flags().StringVarP(&opts.RemoteIpPrefix, "remote-ip-prefix", "r", "", "Remote IP prefix (required)")
	cmdutils.MarkFlagRequired(&cmd, "remote-ip-prefix")

	cmd.Flags().StringVarP(&opts.Protocol, "protocol", "p", "", "Protocol (required)")
	cmdutils.MarkFlagRequired(&cmd, "protocol")
	cmdutils.SetFlagCompletionValues(&cmd, "protocol", types.SecurityGroupProtocols.Keys()...)

	cmd.Flags().Int32Var(&opts.MinPort, "min-port", -1, "Port range minimum")
	cmd.Flags().Int32Var(&opts.MaxPort, "max-port", -1, "Port range maximum")

	cmdutils.AddOutputOnlyIDFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, addFields)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	body := models.CreateSecurityGroupCommand{
		Name:                opts.Name,
		PortMaxRange:        opts.MaxPort,
		PortMinRange:        opts.MinPort,
		Protocol:            types.GetSecurityGroupProtocol(opts.Protocol),
		RemoteIPPrefix:      opts.RemoteIpPrefix,
		StandAloneProfileID: opts.StandAloneProfileID,
	}

	params := security_group.NewSecurityGroupCreateParams().WithV(taikungoclient.Version)
	params = params.WithBody(&body)

	response, err := apiClient.Client.SecurityGroup.SecurityGroupCreate(params, apiClient)
	if err == nil {
		return out.PrintResult(response.Payload, addFields)
	}

	return
}
