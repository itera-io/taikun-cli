package add

import (
	"fmt"

	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/checker"
	"github.com/itera-io/taikungoclient/client/ssh_users"
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
			"PUBLIC-KEY", "sshPublicKey",
		),
	},
)

type AddOptions struct {
	AccessProfileID int32
	Name            string
	PublicKey       string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := &cobra.Command{
		Use:   "add <access-profile-id>",
		Short: "Add an SSH user",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.AccessProfileID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}

			isValid, err := sshPublicKeyIsValid(opts.PublicKey)
			if err != nil {
				return
			}
			if !isValid {
				return fmt.Errorf("SSH public key must be valid")
			}

			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Name, "name", "n", "", "Name (required)")
	cmdutils.MarkFlagRequired(cmd, "name")

	cmd.Flags().StringVarP(&opts.PublicKey, "public-key", "p", "", "Public key (required)")
	cmdutils.MarkFlagRequired(cmd, "public-key")

	cmdutils.AddOutputOnlyIDFlag(cmd)
	cmdutils.AddColumnsFlag(cmd, addFields)

	return cmd
}

func sshPublicKeyIsValid(sshPublicKey string) (bool, error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return false, err
	}

	body := models.SSHKeyCommand{
		SSHPublicKey: sshPublicKey,
	}
	params := checker.NewCheckerSSHParams().WithV(api.Version).WithBody(&body)
	_, err = apiClient.Client.Checker.CheckerSSH(params, apiClient)

	return err == nil, nil
}

func addRun(opts *AddOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.CreateSSHUserCommand{
		AccessProfileID: opts.AccessProfileID,
		Name:            opts.Name,
		SSHPublicKey:    opts.PublicKey,
	}

	params := ssh_users.NewSSHUsersCreateParams().WithV(api.Version).WithBody(&body)

	response, err := apiClient.Client.SSHUsers.SSHUsersCreate(params, apiClient)
	if err == nil {
		return out.PrintResult(response.Payload, addFields)
	}

	return
}
