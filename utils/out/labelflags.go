package out

import (
	"fmt"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"strings"
)

func ParseLabelsFlag(labelsData []string) ([]taikuncore.PrometheusLabelListDto, error) {
	labels := make([]taikuncore.PrometheusLabelListDto, len(labelsData))

	for labelIndex, labelData := range labelsData {
		if len(labelData) == 0 {
			return nil, fmt.Errorf("Invalid empty billing rule label")
		}

		tokens := strings.Split(labelData, "=")
		if len(tokens) != 2 {
			return nil, fmt.Errorf("Invalid billing rule label format: %s", labelData)
		}

		labels[labelIndex] = taikuncore.PrometheusLabelListDto{
			Label: *taikuncore.NewNullableString(&tokens[0]),
			Value: *taikuncore.NewNullableString(&tokens[1]),
		}
	}

	return labels, nil

}
