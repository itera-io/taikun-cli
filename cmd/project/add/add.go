package add

import (
	"context"
	"fmt"
	"time"

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
	AccessProfileID          int32
	AlertingProfileID        int32
	AutoUpgrade              bool
	BackupCredentialID       int32
	CloudCredentialID        int32
	DeleteOnExpiration       bool
	ExpirationDate           string
	Flavors                  []string
	KubernetesProfileID      int32
	KubernetesVersion        string
	Monitoring               bool
	Name                     string
	PolicyProfileID          int32
	RouterIDEndRange         int32
	RouterIDStartRange       int32
	TaikunLBFlavor           string
	Cidr                     string
	DeprecatedAutoscalerName string
	AutoscalerMinSize        int32
	AutoscalerMaxSize        int32
	AutoscalerDiskSize       int32
	AutoscalerFlavor         string
	AutoscalerSpot           bool
	SpotFull                 bool
	SpotWorker               bool
	SpotVms                  bool
	SpotMaxPrice             float64
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
	cmd.Flags().Int32VarP(&opts.PolicyProfileID, "policy-profile-id", "p", 0, "Policy profile ID")
	cmd.Flags().Int32Var(&opts.RouterIDStartRange, "router-id-start-range", -1, "Router ID start range (required with OpenStack and Taikun load balancer")
	cmd.Flags().Int32Var(&opts.RouterIDEndRange, "router-id-end-range", -1, "Router ID end range (required with OpenStack and Taikun load balancer")
	cmd.Flags().StringVar(&opts.TaikunLBFlavor, "taikun-lb-flavor", "", "Taikun load balancer flavor(required with OpenStack and Taikun load balancer")
	cmd.Flags().StringVar(&opts.TaikunLBFlavor, "cidr", "", "Cidr IP")

	cmd.Flags().StringVar(&opts.DeprecatedAutoscalerName, "autoscaler-name", "", "The autoscaler name (specify autoscaler name and flavor to enable autoscaler) DEPRECATED")
	//cmd.Flags().BoolP("autoscaler-name", "", false, "Help message for root version")
	_ = cmd.Flags().MarkDeprecated("autoscaler-name", "just specify autoscaler-flavor. Name will always be taikunCA.")
	cmd.Flags().Int32Var(&opts.AutoscalerDiskSize, "autoscaler-disk-size", 30, "The disk size for the autoscaler in GiB [30 to 8192 GiB] (default 30)")
	cmd.Flags().StringVar(&opts.AutoscalerFlavor, "autoscaler-flavor", "", "The autoscaler flavor (specify flavor to enable autoscaler)")
	cmd.Flags().Int32Var(&opts.AutoscalerMinSize, "autoscaler-min-size", 1, "The minimum size for the autoscaler (default 1)")
	cmd.Flags().Int32Var(&opts.AutoscalerMaxSize, "autoscaler-max-size", 1, "The maximum size for the autoscaler (default 1)")
	cmd.Flags().BoolVar(&opts.AutoscalerSpot, "autoscaler-spot", false, "Enable spot flavors for the autoscaler (default false)")

	cmd.Flags().BoolVar(&opts.SpotFull, "spot-full", false, "Enable full spot flavorsKubernetes (worker + controlplane), bool")
	cmd.Flags().BoolVar(&opts.SpotWorker, "spot-worker", false, "Enable spot flavors for Kubernetes workers, bool")
	cmd.Flags().BoolVar(&opts.SpotVms, "spot-vms", false, "Enable spot flavors for standalone VMs, bool")
	cmd.Flags().Float64Var(&opts.SpotMaxPrice, "spot-max-price", -1, "Set maximum price for spots")

	cmdutils.AddOutputOnlyIDFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, addFields)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	myApiClient := tk.NewClient()

	err = setDefaultAddOptions(opts, myApiClient)
	if err != nil {
		return err
	}

	isKubernetes := true
	body := taikuncore.CreateProjectCommand{
		AccessProfileId:         *taikuncore.NewNullableInt32(&opts.AccessProfileID),
		AlertingProfileId:       *taikuncore.NewNullableInt32(&opts.AlertingProfileID),
		CloudCredentialId:       &opts.CloudCredentialID,
		DeleteOnExpiration:      &opts.DeleteOnExpiration,
		Flavors:                 opts.Flavors,
		IsAutoUpgrade:           &opts.AutoUpgrade,
		IsKubernetes:            &isKubernetes, // Looks to be always true in the UI. No API doc cover what this does.
		KubernetesProfileId:     *taikuncore.NewNullableInt32(&opts.KubernetesProfileID),
		IsMonitoringEnabled:     &opts.Monitoring,
		Name:                    *taikuncore.NewNullableString(&opts.Name),
		AllowFullSpotKubernetes: &opts.SpotFull,
		AllowSpotWorkers:        &opts.SpotWorker,
		AllowSpotVMs:            &opts.SpotVms,
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

	// Autosclaer will enable if you specify name and flavor
	if len(opts.AutoscalerFlavor) != 0 {
		body.SetAutoscalingEnabled(true)
		body.SetAutoscalingFlavor(opts.AutoscalerFlavor)
		//body.SetAutoscalingGroupName(opts.AutoscalerName)
		body.SetMinSize(opts.AutoscalerMinSize)
		body.SetMaxSize(opts.AutoscalerMaxSize)
		body.SetDiskSize(types.GiBToB(opts.AutoscalerDiskSize))
		body.SetAutoscalingSpotEnabled(opts.AutoscalerSpot)
	}

	if opts.SpotMaxPrice > 0 {
		body.SetMaxSpotPrice(opts.SpotMaxPrice)
	}

	data, response, err := myApiClient.Client.ProjectsAPI.ProjectsCreate(context.TODO()).CreateProjectCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	return out.PrintResult(data, addFields)

}

func setDefaultAddOptions(opts *AddOptions, client *tk.Client) (err error) {
	// Get organization ID from cloud credential ID
	organizationID, err := organization.GetOrganizationIDFromCloudCredential(opts.CloudCredentialID, client)
	if err != nil {
		return err
	}

	if opts.AccessProfileID == 0 {
		opts.AccessProfileID, err = getDefaultAccessProfileID(organizationID)
		if err != nil {
			return
		}
	}

	if opts.AlertingProfileID == 0 {
		opts.AlertingProfileID, err = getDefaultAlertingProfileID(organizationID)
		if err != nil {
			return
		}
	}

	if opts.KubernetesProfileID == 0 {
		opts.KubernetesProfileID, err = getDefaultKubernetesProfileID(organizationID)
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
