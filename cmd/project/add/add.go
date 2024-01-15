package add

import (
	"context"
	"fmt"
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/organization"
	"github.com/itera-io/taikun-cli/cmd/project/complete"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"time"
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
			"CLOUD-CREDENTIAL", "cloudCredentialName",
		),
		field.NewVisible(
			"ORG", "organizationName",
		),
		field.NewVisible(
			"ORG-ID", "organizationId",
		),
		field.NewVisible(
			"STATUS", "status",
		),
		field.NewVisible(
			"HEALTH", "health",
		),
		field.NewVisible(
			"K8S", "isKubernetes",
		),
		field.NewVisible(
			"K8S-VERSION", "kubernetesCurrentVersion",
		),
		field.NewVisible(
			"KUBESPRAY-VERSION", "kubesprayCurrentVersion",
		),
		field.NewVisible(
			"CLOUD", "cloudType",
		),
		field.NewVisible(
			"QUOTA-ID", "quotaId",
		),
		field.NewVisibleWithToStringFunc(
			"EXPIRES", "expiredAt", out.FormatDateTimeString,
		),
		field.NewVisible(
			"DELETE-ON-EXPIRATION", "deleteOnExpiration",
		),
		field.NewVisibleWithToStringFunc(
			"LOCK", "isLocked", out.FormatLockStatus,
		),
		field.NewHiddenWithToStringFunc(
			"CREATED-AT", "createdAt", out.FormatDateTimeString,
		),
		field.NewHidden(
			"CREATED-BY", "createdBy",
		),
		field.NewHidden(

			"CIDR", "cidr",
		),
		field.NewVisible(
			"AUTOSCALER", "isAutoscalingEnabled",
		),
	},
)

type AddOptions struct {
	AccessProfileID     int32
	AlertingProfileID   int32
	AutoUpgrade         bool
	BackupCredentialID  int32
	CloudCredentialID   int32
	DeleteOnExpiration  bool
	ExpirationDate      string
	Flavors             []string
	KubernetesProfileID int32
	KubernetesVersion   string
	Monitoring          bool
	Name                string
	OrganizationID      int32
	PolicyProfileID     int32
	RouterIDEndRange    int32
	RouterIDStartRange  int32
	TaikunLBFlavor      string
	Cidr                string
	IsAutoscaler        bool
	AutoscalerName      string
	AutoscalerMinSize   int32
	AutoscalerMaxSize   int32
	AutoscalerDiskSize  float64
	AutoscalerFlavor    string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <name>",
		Short: "Add a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]

			if opts.ExpirationDate != "" {
				if !types.StrIsValidDate(opts.ExpirationDate) {
					return cmderr.ErrUnknownDateFormat
				}
			}

			if opts.RouterIDStartRange != -1 {
				if !types.IsInRouterIDRange(opts.RouterIDStartRange) {
					return cmderr.ErrRouterIDInvalidRange
				}
			}

			if opts.RouterIDEndRange != -1 {
				if !types.IsInRouterIDRange(opts.RouterIDEndRange) {
					return cmderr.ErrRouterIDInvalidRange
				}
			}

			if opts.IsAutoscaler {
				if opts.AutoscalerName == "" {
					return cmderr.NoNameAutoscaler
				}
				if opts.AutoscalerFlavor == "" {
					return cmderr.ErrCheckFailure("The autoscaler flavor")
				}
			}

			return addRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.CloudCredentialID, "cloud-credential-id", "c", 0, "Cloud credential ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "cloud-credential-id")

	cmd.Flags().StringVarP(&opts.KubernetesVersion, "kubernetes-version", "v", "", "Kubernetes version")
	cmdutils.SetFlagCompletionFunc(&cmd, "kubernetes-version", complete.KubernetesVersionCompletionFunc)

	cmd.Flags().Int32Var(&opts.AccessProfileID, "access-profile-id", 0, "Access profile ID")
	cmd.Flags().Int32Var(&opts.AlertingProfileID, "alerting-profile-id", 0, "Alerting profile ID")
	cmd.Flags().BoolVarP(&opts.AutoUpgrade, "auto-upgrade", "u", false, "Enable auto upgrade")
	cmd.Flags().Int32VarP(&opts.BackupCredentialID, "backup-credential-id", "b", 0, "Backup credential ID")
	cmd.Flags().BoolVarP(&opts.DeleteOnExpiration, "delete-on-expiration", "d", false, "Delete the project on its expiration date")
	cmd.Flags().StringVarP(&opts.ExpirationDate, "expiration-date", "e", "", fmt.Sprintf("Expiration date in the format: %s", types.ExpectedDateFormat))
	cmd.Flags().StringSliceVarP(&opts.Flavors, "flavors", "f", []string{}, "Bind flavors to the project")
	cmd.Flags().Int32VarP(&opts.KubernetesProfileID, "kubernetes-profile-id", "k", 0, "Kubernetes profile ID")
	cmd.Flags().BoolVarP(&opts.Monitoring, "monitoring", "m", false, "Enable monitoring")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")
	cmd.Flags().Int32VarP(&opts.PolicyProfileID, "policy-profile-id", "p", 0, "Policy profile ID")
	cmd.Flags().Int32Var(&opts.RouterIDStartRange, "router-id-start-range", -1, "Router ID start range (required with OpenStack and Taikun load balancer")
	cmd.Flags().Int32Var(&opts.RouterIDEndRange, "router-id-end-range", -1, "Router ID end range (required with OpenStack and Taikun load balancer")
	cmd.Flags().StringVar(&opts.TaikunLBFlavor, "taikun-lb-flavor", "", "Taikun load balancer flavor(required with OpenStack and Taikun load balancer")
	cmd.Flags().StringVar(&opts.TaikunLBFlavor, "cidr", "", "Cidr IP")

	cmd.Flags().BoolVar(&opts.IsAutoscaler, "autoscaler", false, "Enable autoscaler for the project")
	cmd.Flags().StringVar(&opts.AutoscalerName, "autoscaler-name", "", "The autoscaler name")
	cmd.Flags().Int32Var(&opts.AutoscalerMinSize, "autoscaler-min-size", 1, "The minimum size for the autoscaler")
	cmd.Flags().Int32Var(&opts.AutoscalerMaxSize, "autoscaler-max-size", 1, "The maximum size for the autoscaler")
	cmd.Flags().Float64Var(&opts.AutoscalerDiskSize, "autoscaler-disk-size", 30, "The disk size for the autoscaler in GiB [30 to 8192 GiB]")
	cmd.Flags().StringVar(&opts.AutoscalerFlavor, "autoscaler-flavor", "", "The autoscaler flavor")

	cmdutils.AddOutputOnlyIDFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, addFields)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	err = setDefaultAddOptions(opts)
	if err != nil {
		return err
	}

	isKubernetes := true
	body := taikuncore.CreateProjectCommand{
		AccessProfileId:     *taikuncore.NewNullableInt32(&opts.AccessProfileID),
		AlertingProfileId:   *taikuncore.NewNullableInt32(&opts.AlertingProfileID),
		CloudCredentialId:   &opts.CloudCredentialID,
		DeleteOnExpiration:  &opts.DeleteOnExpiration,
		Flavors:             opts.Flavors,
		IsAutoUpgrade:       &opts.AutoUpgrade,
		IsKubernetes:        &isKubernetes, // Looks to be always true in the UI. No API doc cover what this does.
		KubernetesProfileId: *taikuncore.NewNullableInt32(&opts.KubernetesProfileID),
		IsMonitoringEnabled: &opts.Monitoring,
		Name:                *taikuncore.NewNullableString(&opts.Name),
		OrganizationId:      *taikuncore.NewNullableInt32(&opts.OrganizationID),
		AutoscalingEnabled:  &opts.IsAutoscaler,
	}

	if opts.BackupCredentialID != 0 {
		body.SetIsBackupEnabled(true)
		body.SetS3CredentialId(opts.BackupCredentialID)
	}

	if opts.ExpirationDate != "" {
		expiredAt := types.StrToDateTime(opts.ExpirationDate)
		body.SetExpiredAt(time.Time(expiredAt))
	}

	if opts.KubernetesVersion != "" {
		body.SetKubernetesVersion(opts.KubernetesVersion)
	}

	if opts.PolicyProfileID != 0 {
		body.SetOpaProfileId(opts.PolicyProfileID)
	}

	if opts.RouterIDStartRange != -1 {
		body.SetRouterIdStartRange(opts.RouterIDStartRange)
	}

	if opts.RouterIDEndRange != -1 {
		body.SetRouterIdEndRange(opts.RouterIDEndRange)
	}

	if opts.TaikunLBFlavor != "" {
		body.SetTaikunLBFlavor(opts.TaikunLBFlavor)
	}

	if opts.Cidr != "" {
		body.SetCidr(opts.Cidr)
	}

	if opts.IsAutoscaler {
		body.SetAutoscalingFlavor(opts.AutoscalerFlavor)
		body.SetAutoscalingGroupName(opts.AutoscalerName)
		body.SetMinSize(opts.AutoscalerMinSize)
		body.SetMaxSize(opts.AutoscalerMaxSize)
		body.SetDiskSize(float64(types.GiBToB(int(opts.AutoscalerDiskSize))))
	}

	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.ProjectsAPI.ProjectsCreate(context.TODO()).CreateProjectCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	return out.PrintResult(data, addFields)

}

func setDefaultAddOptions(opts *AddOptions) (err error) {
	if opts.OrganizationID == 0 {
		opts.OrganizationID, err = organization.GetDefaultOrganizationID()
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
			return err
		}
	}

	return
}

func getDefaultAccessProfileID(organizationID int32) (id int32, err error) {
	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.AccessProfilesAPI.AccessprofilesList(context.TODO()).OrganizationId(organizationID).Execute()
	if err != nil {
		err = tk.CreateError(response, err)
		return
	}

	for _, profile := range data.Data {
		if profile.GetName() == api.DefaultAccessProfileName {
			id = profile.GetId()
			return
		}
	}

	return
}

func getDefaultAlertingProfileID(organizationID int32) (id int32, err error) {
	myApiclient := tk.NewClient()
	data, response, err := myApiclient.Client.AlertingProfilesAPI.AlertingprofilesList(context.TODO()).OrganizationId(organizationID).Execute()
	if err != nil {
		err = tk.CreateError(response, err)
		return
	}

	for _, profile := range data.Data {
		if profile.GetName() == api.DefaultAlertingProfileName {
			id = profile.GetId()
			return
		}
	}

	return
}

func getDefaultKubernetesProfileID(organizationID int32) (id int32, err error) {
	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.KubernetesProfilesAPI.KubernetesprofilesList(context.TODO()).OrganizationId(organizationID).Execute()
	if err != nil {
		err = tk.CreateError(response, err)
		return
	}

	for _, profile := range data.Data {
		if profile.GetName() == api.DefaultKubernetesProfileName {
			id = profile.GetId()
			return
		}
	}

	return
}
