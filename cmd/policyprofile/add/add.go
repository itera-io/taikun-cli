package add

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"

	"github.com/itera-io/taikungoclient/client/opa_profiles"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type AddOptions struct {
	AllowedRepo           []string
	ForbidHTTPIngress     bool
	ForbidNodePort        bool
	ForbidSpecificTags    []string
	IngressWhitelist      []string
	Name                  string
	OrganizationID        int32
	RequireProbe          bool
	UniqueIngresses       bool
	UniqueServiceSelector bool
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := &cobra.Command{
		Use:   "add <name>",
		Short: "Add a policy profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			return addRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")
	cmd.Flags().BoolVar(&opts.ForbidHTTPIngress, "forbid-http-ingress", false, "Requires Ingress resources to be HTTPS only")
	cmd.Flags().BoolVar(&opts.ForbidNodePort, "forbid-node-port", false, "Disallows all Services with type NodePort")
	cmd.Flags().BoolVar(&opts.RequireProbe, "require-probe", false, "Requires Pods to have readiness and liveness probes")
	cmd.Flags().BoolVar(&opts.UniqueIngresses, "unique-ingresses", false, "Requires all Ingress rule hosts to be unique")
	cmd.Flags().BoolVar(&opts.UniqueServiceSelector, "unique-service-selector", false, "Whether services must have globally unique service selectors or not")

	cmd.Flags().StringSliceVar(&opts.AllowedRepo, "allowed-repos", []string{}, "Requires container images to begin with a string from the specified list")
	cmd.Flags().StringSliceVar(&opts.ForbidSpecificTags, "forbidden-tags", []string{}, "Container images must have an image tag different from the ones in the list")
	cmd.Flags().StringSliceVar(&opts.IngressWhitelist, "ingress-whitelist", []string{}, "Requires Ingress to be allowed")

	cmdutils.AddOutputOnlyIDFlag(cmd)

	return cmd
}

func addRun(opts *AddOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := &models.CreateOpaProfileCommand{
		AllowedRepo:           opts.AllowedRepo,
		ForbidHTTPIngress:     opts.ForbidHTTPIngress,
		ForbidNodePort:        opts.ForbidNodePort,
		ForbidSpecificTags:    opts.ForbidSpecificTags,
		IngressWhitelist:      opts.IngressWhitelist,
		Name:                  opts.Name,
		OrganizationID:        opts.OrganizationID,
		RequireProbe:          opts.RequireProbe,
		UniqueIngresses:       opts.UniqueIngresses,
		UniqueServiceSelector: opts.UniqueServiceSelector,
	}

	params := opa_profiles.NewOpaProfilesCreateParams().WithV(api.Version).WithBody(body)
	response, err := apiClient.Client.OpaProfiles.OpaProfilesCreate(params, apiClient)
	if err == nil {
		out.PrintResult(response.Payload,
			"id",
			"name",
			"organizationName",
			"forbidHttpIngress",
			"allowedRepo",
			"forbidNodePort",
			"forbidSpecificTags",
			"ingressWhitelist",
			"requireProbe",
			"uniqueIngresses",
			"uniqueServiceSelector",
		)
	}

	return
}
