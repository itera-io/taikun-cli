package images

import (
	"errors"

	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/offers"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/publishers"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/skus"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/complete"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/cloud_credentials"
	"github.com/itera-io/taikungoclient/client/images"
	"github.com/itera-io/taikungoclient/models"
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
			return imagesRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.AzurePublisher, "azure-publisher", "p", "", "Azure publisher (ignored if cloud type isn't Azure)")
	cmdutils.SetFlagCompletionFunc(&cmd, "azure-publisher", complete.MakeAzurePublisherCompletionFunc())

	cmd.Flags().StringVarP(&opts.AzureOffer, "azure-offer", "o", "", "Azure offer (ignored if cloud type isn't Azure)")
	cmdutils.SetFlagCompletionFunc(&cmd, "azure-offer", complete.MakeAzureOfferCompletionFunc(&opts.AzurePublisher))

	cmd.Flags().StringVarP(&opts.AzureSKU, "azure-sku", "s", "", "Azure SKU (ignored if cloud type isn't Azure)")
	cmdutils.SetFlagCompletionFunc(&cmd, "azure-sku", complete.MakeAzureSKUCompletionFunc(&opts.AzurePublisher, &opts.AzureOffer))

	cmdutils.AddColumnsFlag(&cmd, imagesFields)
	cmdutils.AddLimitFlag(&cmd)

	return &cmd
}

func imagesRun(opts *ImagesOptions) (err error) {
	images, err := getImages(opts)
	if err != nil {
		return err
	}

	out.PrintResults(images, imagesFields)

	return
}

const (
	AWS = iota
	AZURE
	OPENSTACK
)

func getImages(opts *ImagesOptions) (images interface{}, err error) {
	cloudType, err := getCloudType(opts.CloudCredentialID)
	if err != nil {
		return
	}

	switch cloudType {
	case AWS:
		images, err = getAwsImages(opts.CloudCredentialID)
	case AZURE:
		images, err = getAzureImages(opts)
	case OPENSTACK:
		images, err = getOpenstackImages(opts.CloudCredentialID)
	}

	return
}

func getCloudType(cloudCredentialID int32) (cloudType int, err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := cloud_credentials.NewCloudCredentialsDashboardListParams().WithV(api.Version)
	params = params.WithID(&cloudCredentialID)

	response, err := apiClient.Client.CloudCredentials.CloudCredentialsDashboardList(params, apiClient)
	if err == nil {
		if len(response.Payload.Amazon) == 1 {
			cloudType = AWS
		} else if len(response.Payload.Azure) == 1 {
			cloudType = AZURE
		} else if len(response.Payload.Openstack) == 1 {
			cloudType = OPENSTACK
		} else {
			err = cmderr.ResourceNotFoundError("Cloud credential", cloudCredentialID)
		}
	}

	return
}

func getAwsImages(cloudCredentialID int32) (awsImages interface{}, err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := images.NewImagesAwsImagesParams().WithV(api.Version)
	params = params.WithCloudID(cloudCredentialID)

	images := make([]*models.CommonStringBasedDropdownDto, 0)
	for {
		response, err := apiClient.Client.Images.ImagesAwsImages(params, apiClient)
		if err != nil {
			return nil, err
		}
		images = append(images, response.Payload.Data...)
		count := int32(len(images))
		if config.Limit != 0 && count >= config.Limit {
			break
		}
		if count == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&count)
	}

	if config.Limit != 0 && int32(len(images)) > config.Limit {
		images = images[:config.Limit]
	}

	awsImages = images

	return
}

func getOpenstackImages(cloudCredentialID int32) (openStackImages interface{}, err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := images.NewImagesOpenstackImagesParams().WithV(api.Version)
	params = params.WithCloudID(cloudCredentialID)

	images := make([]*models.CommonStringBasedDropdownDto, 0)
	for {
		response, err := apiClient.Client.Images.ImagesOpenstackImages(params, apiClient)
		if err != nil {
			return nil, err
		}
		images = append(images, response.Payload.Data...)
		count := int32(len(images))
		if config.Limit != 0 && count >= config.Limit {
			break
		}
		if count == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&count)
	}

	if config.Limit != 0 && int32(len(images)) > config.Limit {
		images = images[:config.Limit]
	}

	openStackImages = images

	return
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

func getAllAzureImages(opts *ImagesOptions) (azureImages []*models.CommonStringBasedDropdownDto, err error) {
	publishersOptions := publishers.PublishersOptions{CloudCredentialID: opts.CloudCredentialID}
	publishers, err := publishers.ListPublishers(&publishersOptions)
	if err != nil {
		return
	}

	azureImages = make([]*models.CommonStringBasedDropdownDto, 0)
	for _, publisher := range publishers {
		opts.AzurePublisher = publisher
		moreImages, err := getAzureImagesWithPublisher(opts)
		if err != nil {
			return nil, err
		}
		azureImages = append(azureImages, moreImages...)
		if config.Limit != 0 && int32(len(azureImages)) >= config.Limit {
			break
		}
	}

	if config.Limit != 0 && int32(len(azureImages)) > config.Limit {
		azureImages = azureImages[:config.Limit]
	}

	return
}

func getAzureImagesWithPublisher(opts *ImagesOptions) (azureImages []*models.CommonStringBasedDropdownDto, err error) {
	offersOptions := offers.OffersOptions{
		CloudCredentialID: opts.CloudCredentialID,
		Publisher:         opts.AzurePublisher,
	}
	offers, err := offers.ListOffers(&offersOptions)
	if err != nil {
		return
	}

	azureImages = make([]*models.CommonStringBasedDropdownDto, 0)
	for _, offer := range offers {
		opts.AzureOffer = offer
		moreImages, err := getAzureImagesWithOffer(opts)
		if err != nil {
			return nil, err
		}
		azureImages = append(azureImages, moreImages...)
		if config.Limit != 0 && int32(len(azureImages)) >= config.Limit {
			break
		}
	}

	if config.Limit != 0 && int32(len(azureImages)) > config.Limit {
		azureImages = azureImages[:config.Limit]
	}

	return
}

func getAzureImagesWithOffer(opts *ImagesOptions) (azureImages []*models.CommonStringBasedDropdownDto, err error) {
	skusOptions := skus.SKUsOptions{
		CloudCredentialID: opts.CloudCredentialID,
		Publisher:         opts.AzurePublisher,
		Offer:             opts.AzureOffer,
	}
	skus, err := skus.ListSKUs(&skusOptions)
	if err != nil {
		return
	}

	azureImages = make([]*models.CommonStringBasedDropdownDto, 0)
	for _, sku := range skus {
		opts.AzureSKU = sku
		moreImages, err := getAzureImagesWithSKU(opts)
		if err != nil {
			return nil, err
		}
		azureImages = append(azureImages, moreImages...)
		if config.Limit != 0 && int32(len(azureImages)) >= config.Limit {
			break
		}
	}

	if config.Limit != 0 && int32(len(azureImages)) > config.Limit {
		azureImages = azureImages[:config.Limit]
	}

	return
}

func getAzureImagesWithSKU(opts *ImagesOptions) (azureImages []*models.CommonStringBasedDropdownDto, err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := images.NewImagesAzureImagesParams().WithV(api.Version)
	params = params.WithCloudID(opts.CloudCredentialID)
	params = params.WithPublisherName(opts.AzurePublisher)
	params = params.WithOffer(opts.AzureOffer)
	params = params.WithSku(opts.AzureSKU)

	azureImages = make([]*models.CommonStringBasedDropdownDto, 0)
	for {
		response, err := apiClient.Client.Images.ImagesAzureImages(params, apiClient)
		if err != nil {
			return nil, err
		}
		azureImages = append(azureImages, response.Payload.Data...)
		count := int32(len(azureImages))
		if config.Limit != 0 && count >= config.Limit {
			break
		}
		if count == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&count)
	}

	if config.Limit != 0 && int32(len(azureImages)) > config.Limit {
		azureImages = azureImages[:config.Limit]
	}

	return
}
