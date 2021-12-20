package create

import (
	"taikun-cli/api"
	"taikun-cli/apiconfig"
	"taikun-cli/cmd/cmdutils"
	"taikun-cli/utils/format"

	"github.com/itera-io/taikungoclient/client/projects"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type CreateOptions struct {
	AccessProfileID     int32
	AlertingProfileID   int32
	AutoUpgrade         bool
	BackupCredentialID  int32
	CloudCredentialID   int32
	ExpirationDate      string
	Flavors             []string
	KubernetesProfileID int32
	Monitoring          bool
	Name                string
	OrganizationID      int32
	PolicyProfileID     int32
	RouterIDEndRange    int32
	RouterIDStartRange  int32
	TaikunLBFlavor      string
}

func NewCmdCreate() *cobra.Command {
	var opts CreateOptions

	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			return createRun(&opts)
		},
	}

	cmd.Flags().Int32Var(
		&opts.AccessProfileID, "access-profile-id", 0,
		"Access profile ID",
	)

	cmd.Flags().Int32VarP(
		&opts.CloudCredentialID, "cloud-credential-id", "c", 0,
		"Cloud credential ID (required)",
	)
	cmdutils.MarkFlagRequired(cmd, "cloud-credential-id")

	return cmd
}

func createRun(opts *CreateOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.CreateProjectCommand{
		AccessProfileID:   opts.AccessProfileID,
		CloudCredentialID: opts.CloudCredentialID,
		Name:              opts.Name,
	}

	params := projects.NewProjectsCreateParams().WithV(apiconfig.Version)
	params = params.WithBody(&body)

	rep, err := apiClient.Client.Projects.ProjectsCreate(params, apiClient)
	if err == nil {
		format.PrettyPrintJson(rep.Payload)
	}

	return
}
