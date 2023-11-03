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
			return imagesRun(&opts)
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

	cmdutils.AddColumnsFlag(&cmd, imagesFields)
	cmdutils.AddLimitFlag(&cmd, &opts.Limit)

	return &cmd
}

func imagesRun(opts *ImagesOptions) (err error) {
	images, err := getImages(opts)
	if err != nil {
		return err
	}

	return out.PrintResults(images, imagesFields)
}

func getImages(opts *ImagesOptions) (images interface{}, err error) {
	cloudType, err := utils.GetCloudType(opts.CloudCredentialID)
	if err != nil {
		return
	}

	switch cloudType {
	case utils.AWS:
		images, err = getAwsImages(opts)
	case utils.AZURE:
		images, err = getAzureImages(opts)
	case utils.OPENSTACK:
		images, err = getOpenstackImages(opts)
	case utils.GOOGLE:
		images, err = getGoogleImages(opts)
	}

	return
}

func getAwsImages(opts *ImagesOptions) (awsImages interface{}, err error) {
	myApiClient := tk.NewClient()

	// Get owners
	data, response, err := myApiClient.Client.AWSCloudCredentialAPI.AwsOwners(context.TODO()).Execute()
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

	myRequest := myApiClient.Client.ImagesAPI.ImagesAwsImagesList(context.TODO())

	images := make([]taikuncore.AwsExtendedImagesListDto, 0)

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

func getOpenstackImages(opts *ImagesOptions) (openStackImages interface{}, err error) {
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.ImagesAPI.ImagesOpenstackImages(context.TODO(), opts.CloudCredentialID)

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

func getGoogleImages(opts *ImagesOptions) (googleImages interface{}, err error) {
	if opts.GoogleImageType == "" {
		return nil, errors.New(`required flag(s) "google-image-type" not set`)
	}

	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.ImagesAPI.ImagesGoogleImages(context.TODO(), opts.CloudCredentialID, opts.GoogleImageType)

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

func getAzureImages(opts *ImagesOptions) (azureImages interface{}, err error) {
	if opts.AzureSKU != "" {
		if opts.AzureOffer == "" || opts.AzurePublisher == "" {
			return nil, errors.New("before setting --azure-sku, please set --azure-publisher and --azure-offer")
		}

		return getAzureImagesWithSKU(opts)
	}

	if opts.AzureOffer != "" {
		if opts.AzurePublisher == "" {
			return nil, errors.New("before settings --azure-offer, please set --azure-publisher")
		}

		return getAzureImagesWithOffer(opts)
	}

	if opts.AzurePublisher != "" {
		return getAzureImagesWithPublisher(opts)
	}

	return getAllAzureImages(opts)
}

func getAllAzureImages(opts *ImagesOptions) (azureImages []taikuncore.CommonStringBasedDropdownDto, err error) {
	publishersOptions := publishers.PublishersOptions{CloudCredentialID: opts.CloudCredentialID}

	myPublishers, err := publishers.ListPublishers(&publishersOptions)
	if err != nil {
		return nil, err
	}

	azureImages = make([]taikuncore.CommonStringBasedDropdownDto, 0)

	for _, publisher := range myPublishers {
		opts.AzurePublisher = publisher

		moreImages, err := getAzureImagesWithPublisher(opts)
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

func getAzureImagesWithPublisher(opts *ImagesOptions) (azureImages []taikuncore.CommonStringBasedDropdownDto, err error) {
	offersOptions := offers.OffersOptions{
		CloudCredentialID: opts.CloudCredentialID,
		Publisher:         opts.AzurePublisher,
	}

	myOffers, err := offers.ListOffers(&offersOptions)
	if err != nil {
		return nil, err
	}

	azureImages = make([]taikuncore.CommonStringBasedDropdownDto, 0)

	for _, offer := range myOffers {
		opts.AzureOffer = offer

		moreImages, err := getAzureImagesWithOffer(opts)
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

func getAzureImagesWithOffer(opts *ImagesOptions) (azureImages []taikuncore.CommonStringBasedDropdownDto, err error) {
	skusOptions := skus.SKUsOptions{
		CloudCredentialID: opts.CloudCredentialID,
		Publisher:         opts.AzurePublisher,
		Offer:             opts.AzureOffer,
	}

	mySkus, err := skus.ListSKUs(&skusOptions)
	if err != nil {
		return nil, err
	}

	azureImages = make([]taikuncore.CommonStringBasedDropdownDto, 0)

	for _, sku := range mySkus {
		opts.AzureSKU = sku

		moreImages, err := getAzureImagesWithSKU(opts)
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

func getAzureImagesWithSKU(opts *ImagesOptions) (azureImages []taikuncore.CommonStringBasedDropdownDto, err error) {
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.ImagesAPI.ImagesAzureImages(context.TODO(), opts.CloudCredentialID, opts.AzurePublisher, opts.AzureOffer, opts.AzureSKU)
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
