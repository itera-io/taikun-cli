package create

import (
	"fmt"
	"taikun-cli/api"
	"taikun-cli/utils"

	"github.com/itera-io/taikungoclient/client/checker"
	"github.com/itera-io/taikungoclient/client/ssh_users"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type CreateOptions struct {
	AccessProfileID int32
	Name            string
	PublicKey       string
}

func NewCmdCreate() *cobra.Command {
	var opts CreateOptions

	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create SSH user",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]

			isValid, err := sshPublicKeyIsValid(opts.PublicKey)
			if err != nil {
				return err
			}
			if !isValid {
				return fmt.Errorf("SSH public key must be valid")
			}

			return createRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.AccessProfileID, "access-profile-id", "a", 0, "Access profile's ID (required)")
	utils.MarkFlagRequired(cmd, "access-profile-id")

	cmd.Flags().StringVarP(&opts.PublicKey, "public-key", "p", "", "Public key (required)")
	utils.MarkFlagRequired(cmd, "public-key")

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
	params := checker.NewCheckerSSHParams().WithV(utils.ApiVersion).WithBody(&body)
	_, err = apiClient.Client.Checker.CheckerSSH(params, apiClient)

	return err == nil, nil
}

func createRun(opts *CreateOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.CreateSSHUserCommand{
		AccessProfileID: opts.AccessProfileID,
		Name:            opts.Name,
		SSHPublicKey:    opts.PublicKey,
	}

	params := ssh_users.NewSSHUsersCreateParams().WithV(utils.ApiVersion).WithBody(&body)
	response, err := apiClient.Client.SSHUsers.SSHUsersCreate(params, apiClient)
	if err == nil {
		utils.PrettyPrintJson(response)
	}

	return
}
