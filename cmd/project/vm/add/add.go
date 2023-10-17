package add

import (
	"context"
	"errors"
	"fmt"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"strings"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/project/vm/add/complete"
	"github.com/itera-io/taikun-cli/utils/out"
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
			"PROJECT", "projectName",
		),
		field.NewHidden(
			"PROJECT-ID", "projectId",
		),
		field.NewVisible(
			"FLAVOR", "flavorId",
		),
		field.NewVisible(
			"IMAGE", "imageName",
		),
		field.NewVisible(
			"PROFILE", "standAloneProfile/name",
		),
		field.NewHidden(
			"PROFILE-ID", "standAloneProfile/id",
		),
		field.NewVisible(
			"PUBLIC-IP", "publicIpEnabled",
		),
		field.NewVisible(
			"VOLUME-SIZE", "volumeSize",
		),
		field.NewVisible(
			"VOLUME-TYPE", "volumeType",
		),
		field.NewVisibleWithToStringFunc(
			"RAM", "ram", out.FormatBToGiB,
		),
		field.NewVisible(
			"CPU", "cpu",
		),
		field.NewVisible(
			"STATUS", "status",
		),
	},
)

type AddOptions struct {
	CloudInit           string
	Count               int32
	Flavor              string
	ImageID             string
	Name                string
	ProjectID           int32
	PublicIP            bool
	StandAloneProfileID int32
	Tags                []string
	Username            string
	VolumeSize          int64
	VolumeType          string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <project-id>",
		Short: "Add a standalone VM to a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.CloudInit, "cloud-init", "c", "", "Cloud init")

	cmd.Flags().StringVarP(&opts.Flavor, "flavor", "f", "", "Flavor (required)")
	cmdutils.MarkFlagRequired(&cmd, "flavor")

	cmd.Flags().StringVarP(&opts.ImageID, "image-id", "i", "", "Image ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "image-id")

	cmd.Flags().StringVarP(&opts.Name, "name", "n", "", "Name (required)")
	cmdutils.MarkFlagRequired(&cmd, "name")

	cmd.Flags().BoolVar(&opts.PublicIP, "public-ip", false, "Public IP")

	cmd.Flags().Int32VarP(&opts.Count, "count", "x", 1, "Number of VMs to create with this configuration (optional)")

	cmd.Flags().Int32VarP(&opts.StandAloneProfileID, "standalone-profile-id", "s", 0, "Standalone profile ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "standalone-profile-id")

	cmd.Flags().StringSliceVarP(&opts.Tags, "tags", "t", []string{}, `Tags (format: "key=value,key2=value2,...")`)

	cmd.Flags().StringVarP(&opts.Username, "username", "u", "", "Username (optional)")

	cmd.Flags().Int64Var(&opts.VolumeSize, "volume-size", 0, "Volume size in GiB (required)")
	cmdutils.MarkFlagRequired(&cmd, "volume-size")

	cmd.Flags().StringVar(&opts.VolumeType, "volume-type", "", "Volume type")
	cmdutils.SetFlagCompletionFunc(&cmd, "volume-type", complete.VolumeTypeCompletionFunc)

	cmdutils.AddOutputOnlyIDFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, addFields)

	return &cmd
}

func parseTagsOption(tagsOption []string) ([]taikuncore.StandAloneMetaDataDto, error) {
	tags := make([]taikuncore.StandAloneMetaDataDto, len(tagsOption))

	for tagIndex, tag := range tagsOption {
		if len(tag) == 0 {
			return nil, errors.New("Invalid empty VM tag")
		}

		tokens := strings.Split(tag, "=")
		if len(tokens) != 2 {
			return nil, fmt.Errorf("Invalid VM tag format: %s", tag)
		}

		tags[tagIndex] = taikuncore.StandAloneMetaDataDto{
			Key:   *taikuncore.NewNullableString(&tokens[0]),
			Value: *taikuncore.NewNullableString(&tokens[1]),
		}
	}

	return tags, nil
}

func addRun(opts *AddOptions) error {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()
	tags, err := parseTagsOption(opts.Tags)
	if err != nil {
		return err
	}

	// Prepare the arguments for the query
	body := taikuncore.CreateStandAloneVmCommand{
		CloudInit:           *taikuncore.NewNullableString(&opts.CloudInit),
		Count:               &opts.Count,
		FlavorName:          *taikuncore.NewNullableString(&opts.Flavor),
		Image:               *taikuncore.NewNullableString(&opts.ImageID),
		Name:                *taikuncore.NewNullableString(&opts.Name),
		ProjectId:           &opts.ProjectID,
		PublicIpEnabled:     &opts.PublicIP,
		StandAloneMetaDatas: tags,
		StandAloneProfileId: &opts.StandAloneProfileID,
		StandAloneVmDisks:   make([]taikuncore.StandAloneVmDiskDto, 0),
		VolumeSize:          &opts.VolumeSize,
	}

	if opts.Username != "" {
		body.Username = *taikuncore.NewNullableString(&opts.Username)
	}

	if opts.VolumeType != "" {
		body.VolumeType = *taikuncore.NewNullableString(&opts.VolumeType)
	}

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.StandaloneAPI.StandaloneCreate(context.TODO()).CreateStandAloneVmCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	// Manipulate the gathered data
	return out.PrintResult(data, addFields)

	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return err
		}

		tags, err := parseTagsOption(opts.Tags)
		if err != nil {
			return err
		}

		body := models.CreateStandAloneVMCommand{
			CloudInit:           opts.CloudInit,
			Count:               1,
			FlavorName:          opts.Flavor,
			Image:               opts.ImageID,
			Name:                opts.Name,
			ProjectID:           opts.ProjectID,
			PublicIPEnabled:     opts.PublicIP,
			StandAloneMetaDatas: tags,
			StandAloneProfileID: opts.StandAloneProfileID,
			StandAloneVMDisks:   make([]*models.StandAloneVMDiskDto, 0),
			VolumeSize:          opts.VolumeSize,
		}

		if opts.Username != "" {
			body.Username = opts.Username
		}

		if opts.VolumeType != "" {
			body.VolumeType = opts.VolumeType
		}

		params := stand_alone.NewStandAloneCreateParams().WithV(taikungoclient.Version)
		params = params.WithBody(&body)

		response, err := apiClient.Client.StandAlone.StandAloneCreate(params, apiClient)
		if err != nil {
			return err
		}

		return out.PrintResult(response.Payload, addFields)
	*/
}
