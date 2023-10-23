package types

import (
	"github.com/itera-io/taikun-cli/utils/gmap"
	taikuncore "github.com/itera-io/taikungoclient/client"
)

var EPrometheusTypes = gmap.New(
	map[string]interface{}{
		"count": taikuncore.PROMETHEUSTYPE_COUNT,
		"sum":   taikuncore.PROMETHEUSTYPE_SUM,
	},
)

var PrometheusTypes = gmap.New(
	map[string]interface{}{
		"count": taikuncore.PROMETHEUSTYPE_COUNT,
		"sum":   taikuncore.PROMETHEUSTYPE_SUM,
	},
)

//func GetEPrometheusType(showbackType string) taikuncore.prometheustype {
//	model, _ := EPrometheusTypes.Get(showbackType).(taikuncore.PrometheusType)
//	return model
//}

func GetPrometheusType(showbackType string) taikuncore.PrometheusType {
	model, _ := PrometheusTypes.Get(showbackType).(taikuncore.PrometheusType)
	return model
}

//import (
//	"github.com/itera-io/taikun-cli/utils/gmap"
//	"github.com/itera-io/taikungoclient/models"
//)
//
//var EPrometheusTypes = gmap.New(
//	map[string]interface{}{
//		"count": models.EPrometheusType(100),
//		"sum":   models.EPrometheusType(200),
//	},
//)
//
//var PrometheusTypes = gmap.New(
//	map[string]interface{}{
//		"count": models.PrometheusType(100),
//		"sum":   models.PrometheusType(200),
//	},
//)
//
//func GetEPrometheusType(showbackType string) models.EPrometheusType {
//	model, _ := EPrometheusTypes.Get(showbackType).(models.EPrometheusType)
//	return model
//}
//
//func GetPrometheusType(showbackType string) models.PrometheusType {
//	model, _ := PrometheusTypes.Get(showbackType).(models.PrometheusType)
//	return model
//}
