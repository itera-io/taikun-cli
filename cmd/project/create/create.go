package create

import (
	"taikun-cli/api"
	"taikun-cli/apiconfig"
	"taikun-cli/cmd/cmdutils"
	"taikun-cli/utils/format"

	"github.com/itera-io/taikungoclient/client/access_profiles"
	"github.com/itera-io/taikungoclient/client/alerting_profiles"
	"github.com/itera-io/taikungoclient/client/projects"
	"github.com/itera-io/taikungoclient/client/users"
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

	cmd.Flags().Int32VarP(
		&opts.CloudCredentialID, "cloud-credential-id", "c", 0,
		"Cloud credential ID (required)",
	)
	cmdutils.MarkFlagRequired(cmd, "cloud-credential-id")

	cmd.Flags().Int32Var(
		&opts.AccessProfileID, "access-profile-id", 0,
		"Access profile ID",
	)

	cmd.Flags().Int32Var(
		&opts.AlertingProfileID, "alerting-profile-id", 0,
		"Alerting profile ID",
	)

	cmd.Flags().Int32VarP(
		&opts.OrganizationID, "organization-id", "o", 0,
		"Organization ID",
	)

	return cmd
}

func createRun(opts *CreateOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	if opts.OrganizationID == 0 {
		opts.OrganizationID, err = getDefaultOrganizationID()
		if err != nil {
			return
		}
	}
	if opts.AccessProfileID == 0 {
		opts.AccessProfileID, err = getDefaultAccessProfileID(opts.OrganizationID)
		if err != nil {
			return
		}
	}
	if opts.AlertingProfileID == 0 {
		opts.AlertingProfileID, err = getDefaultAlertingProfileID(opts.OrganizationID)
		if err != nil {
			return
		}
	}

	body := models.CreateProjectCommand{
		AccessProfileID:   opts.AccessProfileID,
		AlertingProfileID: opts.AlertingProfileID,
		CloudCredentialID: opts.CloudCredentialID,
		Name:              opts.Name,
		OrganizationID:    opts.OrganizationID,
	}

	params := projects.NewProjectsCreateParams().WithV(apiconfig.Version)
	params = params.WithBody(&body)

	response, err := apiClient.Client.Projects.ProjectsCreate(params, apiClient)
	if err == nil {
		format.PrettyPrintJson(response.Payload)
	}

	return
}

func getDefaultOrganizationID() (id int32, err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}
	params := users.NewUsersDetailsParams().WithV(apiconfig.Version)
	response, err := apiClient.Client.Users.UsersDetails(params, apiClient)
	if err == nil {
		id = response.Payload.Data.OrganizationID
	}
	return
}

func getDefaultAccessProfileID(organizationID int32) (id int32, err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := access_profiles.NewAccessProfilesAccessProfilesForOrganizationListParams()
	params = params.WithV(apiconfig.Version).WithOrganizationID(&organizationID)
	response, err := apiClient.Client.AccessProfiles.AccessProfilesAccessProfilesForOrganizationList(params, apiClient)
	if err != nil {
		return
	}

	for _, profile := range response.Payload {
		if profile.Name == apiconfig.DefaultAccessProfileName {
			id = profile.ID
			return
		}
	}
	return
}

func getDefaultAlertingProfileID(organizationID int32) (id int32, err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := alerting_profiles.NewAlertingProfilesAlertingProfilesForOrganizationListParams()
	params = params.WithV(apiconfig.Version).WithOrganizationID(&organizationID)
	response, err := apiClient.Client.AlertingProfiles.AlertingProfilesAlertingProfilesForOrganizationList(params, apiClient)
	if err != nil {
		return
	}

	for _, profile := range response.Payload {
		if profile.Name == apiconfig.DefaultAlertingProfileName {
			id = profile.ID
			return
		}
	}
	return
}
