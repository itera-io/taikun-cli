package add

import (
	"context"
	"fmt"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
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
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.SshKeyCommand{
		SshPublicKey: *taikuncore.NewNullableString(&sshPublicKey),
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.CheckerAPI.CheckerSsh(context.TODO()).SshKeyCommand(body).Execute()
	if err != nil {
		return false, tk.CreateError(response, err)
	}
	return err == nil, nil
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return false, err
		}

		body := models.SSHKeyCommand{
			SSHPublicKey: sshPublicKey,
		}
		params := checker.NewCheckerSSHParams().WithV(taikungoclient.Version).WithBody(&body)
		_, err = apiClient.Client.Checker.CheckerSSH(params, apiClient)

		return err == nil, nil
	*/
}

func addRun(opts *AddOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.CreateSshUserCommand{
		Name:            *taikuncore.NewNullableString(&opts.Name),
		SshPublicKey:    *taikuncore.NewNullableString(&opts.PublicKey),
		AccessProfileId: &opts.AccessProfileID,
	}

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.SshUsersAPI.SshusersCreate(context.TODO()).CreateSshUserCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	return out.PrintResult(data, addFields)
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		body := models.CreateSSHUserCommand{
			AccessProfileID: opts.AccessProfileID,
			Name:            opts.Name,
			SSHPublicKey:    opts.PublicKey,
		}

		params := ssh_users.NewSSHUsersCreateParams().WithV(taikungoclient.Version).WithBody(&body)

		response, err := apiClient.Client.SSHUsers.SSHUsersCreate(params, apiClient)
		if err == nil {
			return out.PrintResult(response.Payload, addFields)
		}

		return
	*/
}
