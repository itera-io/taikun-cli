package images

import (
	"context"
	"errors"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"

	"github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/offers"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/publishers"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/skus"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/complete"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/utils"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/spf13/cobra"
)

var imagesFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"NAME", "name",
		),
	},
)

type ImagesOptions struct {
	CloudCredentialID int32
	AzurePublisher    string
	AzureOffer        string
	AzureSKU          string
	GoogleImageType   string
	GoogleLatest      bool
	Limit             int32
}

func NewCmdImages() *cobra.Command {
	var opts ImagesOptions

	cmd := cobra.Command{
		Use:   "images <cloud-credential-id>",
		Short: "List a cloud credential's available images",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.CloudCredentialID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			if opts.GoogleImageType != "" {
				if err := cmdutils.CheckFlagValue("google-image-type", opts.GoogleImageType, types.GoogleImageTypes); err != nil {
					return err
				}
			}
			return imagesRun(cmd, &opts)
		},
	}

	cmd.Flags().StringVarP(&opts.AzurePublisher, "azure-publisher", "p", "", "Azure publisher (ignored if cloud type isn't Azure)")
	cmdutils.SetFlagCompletionFunc(&cmd, "azure-publisher", complete.MakeAzurePublisherCompletionFunc())

	cmd.Flags().StringVarP(&opts.AzureOffer, "azure-offer", "o", "", "Azure offer (ignored if cloud type isn't Azure)")
	cmdutils.SetFlagCompletionFunc(&cmd, "azure-offer", complete.MakeAzureOfferCompletionFunc(&opts.AzurePublisher))

	cmd.Flags().StringVarP(&opts.AzureSKU, "azure-sku", "s", "", "Azure SKU (ignored if cloud type isn't Azure)")
	cmdutils.SetFlagCompletionFunc(&cmd, "azure-sku", complete.MakeAzureSKUCompletionFunc(&opts.AzurePublisher, &opts.AzureOffer))

	cmd.Flags().StringVarP(&opts.GoogleImageType, "google-image-type", "g", "", "Google image type (ignored if cloud type isn't Google)")
	cmdutils.SetFlagCompletionValues(&cmd, "google-image-type", types.GoogleImageTypes.Keys()...)

	cmd.Flags().BoolVarP(&opts.GoogleLatest, "google-latest", "l", false, "Google flag for latest images (ignored if cloud type isn't Google)")

	cmdutils.AddColumnsFlag(&cmd, imagesFields)
	cmdutils.AddLimitFlag(&cmd, &opts.Limit)

	return &cmd
}

func imagesRun(cmd *cobra.Command, opts *ImagesOptions) (err error) {
	images, err := getImages(cmd, opts)
	if err != nil {
		return err
	}

	return out.PrintResults(images, imagesFields)
}

func getImages(cmd *cobra.Command, opts *ImagesOptions) (images interface{}, err error) {
	cloudType, err := utils.GetCloudType(cmd, opts.CloudCredentialID)
	if err != nil {
		return
	}

	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	switch cloudType {
	case taikuncore.CLOUDTYPE_AWS:
		images, err = getAwsImages(ctx, opts)
	case taikuncore.CLOUDTYPE_AZURE:
		images, err = getAzureImages(cmd, ctx, opts)
	case taikuncore.CLOUDTYPE_OPENSTACK:
		images, err = getOpenstackImages(ctx, opts)
	case taikuncore.CLOUDTYPE_GOOGLE:
		images, err = getGoogleImages(ctx, opts)
	case taikuncore.CLOUDTYPE_PROXMOX:
		images, err = getProxmoxImages(ctx, opts)
	case taikuncore.CLOUDTYPE_VSPHERE:
		images, err = getVsphereImages(ctx, opts)
	}

	return
}

func getVsphereImages(ctx context.Context, opts *ImagesOptions) (vsphereImages interface{}, err error) {
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.ImagesAPI.ImagesVsphereImages(ctx, opts.CloudCredentialID)

	images := make([]taikuncore.CommonStringBasedDropdownDto, 0)

	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			err = tk.CreateError(response, err)
			return nil, err
		}

		images = append(images, data.GetData()...)

		count := int32(len(images))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}

		if count == data.GetTotalCount() {
			break
		}

		myRequest = myRequest.Offset(count)
	}

	if opts.Limit != 0 && int32(len(images)) > opts.Limit {
		images = images[:opts.Limit]
	}
	vsphereImages = images

	return vsphereImages, nil
}

func getProxmoxImages(ctx context.Context, opts *ImagesOptions) (proxmoxImages interface{}, err error) {
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.ImagesAPI.ImagesProxmoxImages(ctx, opts.CloudCredentialID)

	images := make([]taikuncore.CommonStringBasedDropdownDto, 0)

	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			err = tk.CreateError(response, err)
			return nil, err
		}

		images = append(images, data.GetData()...)

		count := int32(len(images))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}

		if count == data.GetTotalCount() {
			break
		}

		myRequest = myRequest.Offset(count)
	}

	if opts.Limit != 0 && int32(len(images)) > opts.Limit {
		images = images[:opts.Limit]
	}

	proxmoxImages = images

	return proxmoxImages, nil
}

func getAwsImages(ctx context.Context, opts *ImagesOptions) (awsImages interface{}, err error) {
	myApiClient := tk.NewClient()

	// Get owners
	data, response, err := myApiClient.Client.AWSCloudCredentialAPI.AwsOwners(ctx).Execute()
	if err != nil {
		return nil, tk.CreateError(response, err)
	}
	owners := make([]string, 0)
	for i := 0; i < len(data); i++ {
		owners = append(owners, data[i].GetId())
	}

	// Get images
	body := taikuncore.AwsImagesPostListCommand{
		CloudId: &opts.CloudCredentialID,
		Owners:  owners,
	}

	myRequest := myApiClient.Client.ImagesAPI.ImagesAwsImagesList(ctx)

	images := make([]taikuncore.CommonStringBasedDropdownDto, 0)

	for {
		data, response, responseErr := myRequest.AwsImagesPostListCommand(body).Execute()
		if responseErr != nil {
			err = tk.CreateError(response, responseErr)
			return
		}

		images = append(images, data.GetData()...)

		count := int32(len(images))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}

		if count == data.GetTotalCount() {
			break
		}

		body.SetOffset(count)
	}

	if opts.Limit != 0 && int32(len(images)) > opts.Limit {
		images = images[:opts.Limit]
	}

	awsImages = images

	return awsImages, nil

}

func getOpenstackImages(ctx context.Context, opts *ImagesOptions) (openStackImages interface{}, err error) {
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.ImagesAPI.ImagesOpenstackImages(ctx, opts.CloudCredentialID)

	images := make([]taikuncore.CommonStringBasedDropdownDto, 0)

	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			err = tk.CreateError(response, err)
			return nil, err
		}

		images = append(images, data.GetData()...)

		count := int32(len(images))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}

		if count == data.GetTotalCount() {
			break
		}

		myRequest = myRequest.Offset(count)
	}

	if opts.Limit != 0 && int32(len(images)) > opts.Limit {
		images = images[:opts.Limit]
	}

	openStackImages = images

	return openStackImages, nil

}

func getGoogleImages(ctx context.Context, opts *ImagesOptions) (googleImages interface{}, err error) {
	if opts.GoogleImageType == "" {
		return nil, errors.New(`required flag(s) "google-image-type" not set`)
	}

	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.ImagesAPI.ImagesGoogleImages(ctx, opts.CloudCredentialID, opts.GoogleImageType).Latest(opts.GoogleLatest)

	images := make([]taikuncore.CommonStringBasedDropdownDto, 0)

	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			err = tk.CreateError(response, err)
			return nil, err
		}

		images = append(images, data.GetData()...)

		count := int32(len(images))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}

		if count == data.GetTotalCount() {
			break
		}

		myRequest = myRequest.Offset(count)
	}

	if opts.Limit != 0 && int32(len(images)) > opts.Limit {
		images = images[:opts.Limit]
	}

	googleImages = images

	return googleImages, nil

}

func getAzureImages(cmd *cobra.Command, ctx context.Context, opts *ImagesOptions) (azureImages interface{}, err error) {
	if opts.AzureSKU != "" {
		if opts.AzureOffer == "" || opts.AzurePublisher == "" {
			return nil, errors.New("before setting --azure-sku, please set --azure-publisher and --azure-offer")
		}

		return getAzureImagesWithSKU(cmd, ctx, opts)
	}

	if opts.AzureOffer != "" {
		if opts.AzurePublisher == "" {
			return nil, errors.New("before settings --azure-offer, please set --azure-publisher")
		}

		return getAzureImagesWithOffer(cmd, ctx, opts)
	}

	if opts.AzurePublisher != "" {
		return getAzureImagesWithPublisher(cmd, ctx, opts)
	}

	return getAllAzureImages(cmd, ctx, opts)
}

func getAllAzureImages(cmd *cobra.Command, ctx context.Context, opts *ImagesOptions) (azureImages []taikuncore.CommonStringBasedDropdownDto, err error) {
	publishersOptions := publishers.PublishersOptions{CloudCredentialID: opts.CloudCredentialID}

	myPublishers, err := publishers.ListPublishers(cmd, &publishersOptions)
	if err != nil {
		return nil, err
	}

	azureImages = make([]taikuncore.CommonStringBasedDropdownDto, 0)

	for _, publisher := range myPublishers {
		opts.AzurePublisher = publisher

		moreImages, err := getAzureImagesWithPublisher(cmd, ctx, opts)
		if err != nil {
			return nil, err
		}

		azureImages = append(azureImages, moreImages...)
		if opts.Limit != 0 && int32(len(azureImages)) >= opts.Limit {
			break
		}
	}

	if opts.Limit != 0 && int32(len(azureImages)) > opts.Limit {
		azureImages = azureImages[:opts.Limit]
	}

	return azureImages, nil
}

func getAzureImagesWithPublisher(cmd *cobra.Command, ctx context.Context, opts *ImagesOptions) (azureImages []taikuncore.CommonStringBasedDropdownDto, err error) {
	offersOptions := offers.OffersOptions{
		CloudCredentialID: opts.CloudCredentialID,
		Publisher:         opts.AzurePublisher,
	}

	myOffers, err := offers.ListOffers(cmd, &offersOptions)
	if err != nil {
		return nil, err
	}

	azureImages = make([]taikuncore.CommonStringBasedDropdownDto, 0)

	for _, offer := range myOffers {
		opts.AzureOffer = offer

		moreImages, err := getAzureImagesWithOffer(cmd, ctx, opts)
		if err != nil {
			return nil, err
		}

		azureImages = append(azureImages, moreImages...)
		if opts.Limit != 0 && int32(len(azureImages)) >= opts.Limit {
			break
		}
	}

	if opts.Limit != 0 && int32(len(azureImages)) > opts.Limit {
		azureImages = azureImages[:opts.Limit]
	}

	return azureImages, nil
}

func getAzureImagesWithOffer(cmd *cobra.Command, ctx context.Context, opts *ImagesOptions) (azureImages []taikuncore.CommonStringBasedDropdownDto, err error) {
	skusOptions := skus.SKUsOptions{
		CloudCredentialID: opts.CloudCredentialID,
		Publisher:         opts.AzurePublisher,
		Offer:             opts.AzureOffer,
	}

	mySkus, err := skus.ListSKUs(cmd, &skusOptions)
	if err != nil {
		return nil, err
	}

	azureImages = make([]taikuncore.CommonStringBasedDropdownDto, 0)

	for _, sku := range mySkus {
		opts.AzureSKU = sku

		moreImages, err := getAzureImagesWithSKU(cmd, ctx, opts)
		if err != nil {
			return nil, err
		}

		azureImages = append(azureImages, moreImages...)
		if opts.Limit != 0 && int32(len(azureImages)) >= opts.Limit {
			break
		}
	}

	if opts.Limit != 0 && int32(len(azureImages)) > opts.Limit {
		azureImages = azureImages[:opts.Limit]
	}

	return azureImages, nil
}

func getAzureImagesWithSKU(cmd *cobra.Command, ctx context.Context, opts *ImagesOptions) (azureImages []taikuncore.CommonStringBasedDropdownDto, err error) {
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.ImagesAPI.ImagesAzureImages(ctx, opts.CloudCredentialID, opts.AzurePublisher, opts.AzureOffer, opts.AzureSKU)
	azureImages = make([]taikuncore.CommonStringBasedDropdownDto, 0)

	for {
		data, response, responseErr := myRequest.Execute()
		if responseErr != nil {
			err = tk.CreateError(response, responseErr)
			return
		}

		azureImages = append(azureImages, data.GetData()...)

		count := int32(len(azureImages))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}

		if count == data.GetTotalCount() {
			break
		}

		myRequest = myRequest.Offset(count)
	}

	if opts.Limit != 0 && int32(len(azureImages)) > opts.Limit {
		azureImages = azureImages[:opts.Limit]
	}

	return azureImages, nil

}
