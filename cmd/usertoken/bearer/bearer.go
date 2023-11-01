package bearer

import (
	"context"
	"fmt"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

var ListFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"TOKEN", "token",
		),
	},
)

type ListOptions struct {
	decorate bool
}

func NewCmdBearer() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "get-bearer",
		Short: "Get a bearer token for manual API testing",
		RunE: func(cmd *cobra.Command, args []string) error {
			return bearerRun(&opts)
		},
		Args:    cobra.NoArgs,
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().BoolVarP(&opts.decorate, "decorate", "d", false, "Enable to use decoration table, not just plain token.")

	return cmd
}

func bearerRun(opts *ListOptions) (err error) {
	// Connect to the API and retrieve data.
	myApiClient := tk.NewClient()
	_, response, err := myApiClient.Client.UserTokenAPI.UsertokenList(context.TODO()).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	plainToken := fmt.Sprint("Bearer " + myApiClient.GetToken())

	// Return just plain token
	if !opts.decorate && !config.Quiet {
		fmt.Println(plainToken)
		return
	}

	// Return with decorations
	myToken := map[string]string{
		"token": plainToken,
	}
	return out.PrintResult(myToken, ListFields)
}
