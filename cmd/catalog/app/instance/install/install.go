package install

import (
	"context"
	"encoding/base64"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"log"
	"os"
)

type InstallOptions struct {
	Name               string
	ProjectId          int32
	CatalogAppId       int32
	Namespace          string
	ExtraValuesLiteral string
	ExtraValuesFile    string
	Autosync           bool
}

func NewCmdInstall() *cobra.Command {
	var opts InstallOptions

	cmd := cobra.Command{
		Use:   "install <NAME_FOR_APP_INSTANCE> <CATALOG_APP_ID> <PROJECT_ID>",
		Short: "Install catalog app into a project. This will create an app instance.",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.Name = args[0]
			opts.CatalogAppId, err = types.Atoi32(args[1])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			opts.ProjectId, err = types.Atoi32(args[2])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return installAppRun(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.ExtraValuesLiteral, "extra-values-literal", "v", "", "Extra values to pass to the helm chart")
	cmd.Flags().StringVarP(&opts.ExtraValuesFile, "extra-values-file", "f", "", "Extra values to pass to the helm chart")
	cmd.MarkFlagsMutuallyExclusive("extra-values-literal", "extra-values-file")

	cmd.Flags().StringVarP(&opts.Namespace, "namespace", "n", "default", "Namespace to install catalog app into")
	cmd.Flags().BoolVarP(&opts.Autosync, "auto-sync", "", false, "Automatically sync the catalog app together")

	return &cmd
}

func installAppRun(opts InstallOptions) (err error) {
	myApiClient := tk.NewClient()
	var extraValuesB64 string

	switch {
	case opts.ExtraValuesLiteral != "":
		extraValuesB64 = base64.StdEncoding.EncodeToString([]byte(opts.ExtraValuesLiteral))
	case opts.ExtraValuesFile != "":
		content, err := os.ReadFile(opts.ExtraValuesFile)
		if err != nil {
			log.Fatalf("Failed to read file: %v", err)
		}
		extraValuesB64 = base64.StdEncoding.EncodeToString(content)
	default:
		log.Fatal("Neither ExtraValuesLiteral nor ExtraValuesFile is set")
	}

	body := taikuncore.CreateProjectAppCommand{
		Name:         *taikuncore.NewNullableString(&opts.Name),
		Namespace:    *taikuncore.NewNullableString(&opts.Namespace),
		ProjectId:    &opts.ProjectId,
		CatalogAppId: &opts.CatalogAppId,
		ExtraValues:  *taikuncore.NewNullableString(&extraValuesB64),
		AutoSync:     &opts.Autosync,
	}

	_, response, err := myApiClient.Client.ProjectAppsAPI.ProjectappInstall(context.TODO()).CreateProjectAppCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()

	return nil
}
