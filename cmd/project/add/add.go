package add

import (
	"fmt"

	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/access_profiles"
	"github.com/itera-io/taikungoclient/client/alerting_profiles"
	"github.com/itera-io/taikungoclient/client/kubernetes_profiles"
	"github.com/itera-io/taikungoclient/client/projects"
	"github.com/itera-io/taikungoclient/client/users"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type AddOptions struct {
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

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := &cobra.Command{
		Use:   "add <name>",
		Short: "Add a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]

			if opts.ExpirationDate != "" {
				if !types.StrIsValidDate(opts.ExpirationDate) {
					return cmderr.InvalidDateFormatError
				}
			}

			if opts.RouterIDStartRange != -1 {
				if !types.IsInRouterIDRange(opts.RouterIDStartRange) {
					return cmderr.RouterIDRangeError
				}
			}

			if opts.RouterIDEndRange != -1 {
				if !types.IsInRouterIDRange(opts.RouterIDEndRange) {
					return cmderr.RouterIDRangeError
				}
			}

			return addRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.CloudCredentialID, "cloud-credential-id", "c", 0, "Cloud credential ID (required)")
	cmdutils.MarkFlagRequired(cmd, "cloud-credential-id")

	cmd.Flags().Int32Var(&opts.AccessProfileID, "access-profile-id", 0, "Access profile ID")
	cmd.Flags().Int32Var(&opts.AlertingProfileID, "alerting-profile-id", 0, "Alerting profile ID")
	cmd.Flags().BoolVarP(&opts.AutoUpgrade, "auto-upgrade", "u", false, "Enable auto upgrade")
	cmd.Flags().Int32VarP(&opts.BackupCredentialID, "backup-credential-id", "b", 0, "Backup credential ID")
	cmd.Flags().StringVarP(&opts.ExpirationDate, "expiration-date", "e", "", fmt.Sprintf("Expiration date in the format: %s", types.ExpectedDateFormat))
	cmd.Flags().StringSliceVarP(&opts.Flavors, "flavors", "f", []string{}, "Bind flavors to the project")
	cmd.Flags().Int32VarP(&opts.KubernetesProfileID, "kubernetes-profile-id", "k", 0, "Kubernetes profile ID")
	cmd.Flags().BoolVarP(&opts.Monitoring, "monitoring", "m", false, "Enable monitoring")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")
	cmd.Flags().Int32VarP(&opts.PolicyProfileID, "policy-profile-id", "p", 0, "Policy profile ID")
	cmd.Flags().Int32Var(&opts.RouterIDStartRange, "router-id-start-range", -1, "Router ID start range (required with OpenStack and Taikun load balancer")
	cmd.Flags().Int32Var(&opts.RouterIDEndRange, "router-id-end-range", -1, "Router ID end range (required with OpenStack and Taikun load balancer")
	cmd.Flags().StringVar(&opts.TaikunLBFlavor, "taikun-lb-flavor", "", "Taikun load balancer flavor(required with OpenStack and Taikun load balancer")

	cmdutils.AddOutputOnlyIDFlag(cmd)

	return cmd
}

func addRun(opts *AddOptions) (err error) {
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
	if opts.KubernetesProfileID == 0 {
		opts.KubernetesProfileID, err = getDefaultKubernetesProfileID(opts.OrganizationID)
		if err != nil {
			return
		}
	}

	body := models.CreateProjectCommand{
		AccessProfileID:     opts.AccessProfileID,
		AlertingProfileID:   opts.AlertingProfileID,
		CloudCredentialID:   opts.CloudCredentialID,
		Flavors:             opts.Flavors,
		IsAutoUpgrade:       opts.AutoUpgrade,
		IsKubernetes:        true,
		KubernetesProfileID: opts.KubernetesProfileID,
		IsMonitoringEnabled: opts.Monitoring,
		Name:                opts.Name,
		OrganizationID:      opts.OrganizationID,
	}

	if opts.BackupCredentialID != 0 {
		body.IsBackupEnabled = true
		body.S3CredentialID = opts.BackupCredentialID
	}

	if opts.ExpirationDate != "" {
		expiredAt := types.StrToDateTime(opts.ExpirationDate)
		body.ExpiredAt = &expiredAt
	}

	if opts.PolicyProfileID != 0 {
		body.OpaProfileID = opts.PolicyProfileID
	}

	if opts.RouterIDStartRange != -1 {
		body.RouterIDStartRange = opts.RouterIDStartRange
	}
	if opts.RouterIDEndRange != -1 {
		body.RouterIDEndRange = opts.RouterIDEndRange
	}
	if opts.TaikunLBFlavor != "" {
		body.TaikunLBFlavor = opts.TaikunLBFlavor
	}

	params := projects.NewProjectsCreateParams().WithV(api.Version)
	params = params.WithBody(&body)

	response, err := apiClient.Client.Projects.ProjectsCreate(params, apiClient)
	if err == nil {
		out.PrintResult(response.Payload,
			"id",
			"name",
			"organizationName",
			"status",
			"health",
			"createdAt",
			"kubernetesCurrentVersion",
			"cloudType",
			"hasKubeConfigFile",
			"quotaId",
			"expiredAt",
			"isLocked",
		)
	}

	return
}

func getDefaultOrganizationID() (id int32, err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}
	params := users.NewUsersDetailsParams().WithV(api.Version)
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
	params = params.WithV(api.Version).WithOrganizationID(&organizationID)
	response, err := apiClient.Client.AccessProfiles.AccessProfilesAccessProfilesForOrganizationList(params, apiClient)
	if err != nil {
		return
	}

	for _, profile := range response.Payload {
		if profile.Name == api.DefaultAccessProfileName {
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
	params = params.WithV(api.Version).WithOrganizationID(&organizationID)
	response, err := apiClient.Client.AlertingProfiles.AlertingProfilesAlertingProfilesForOrganizationList(params, apiClient)
	if err != nil {
		return
	}

	for _, profile := range response.Payload {
		if profile.Name == api.DefaultAlertingProfileName {
			id = profile.ID
			return
		}
	}
	return
}

func getDefaultKubernetesProfileID(organizationID int32) (id int32, err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := kubernetes_profiles.NewKubernetesProfilesKubernetesProfilesForOrganizationListParams()
	params = params.WithV(api.Version).WithOrganizationID(&organizationID)
	response, err := apiClient.Client.KubernetesProfiles.KubernetesProfilesKubernetesProfilesForOrganizationList(params, apiClient)
	if err != nil {
		return
	}

	for _, profile := range response.Payload {
		if profile.Name == api.DefaultKubernetesProfileName {
			id = profile.ID
			return
		}
	}
	return
}
