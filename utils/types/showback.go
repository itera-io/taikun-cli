package types

import (
	"github.com/itera-io/taikun-cli/utils/gmap"
	taikunshowback "github.com/itera-io/taikungoclient/showbackclient"
)

var ShowbackKinds = gmap.New(
	map[string]interface{}{
		"general":  taikunshowback.ESHOWBACKTYPE_GENERAL,
		"external": taikunshowback.ESHOWBACKTYPE_EXTERNAL,
	},
)

func GetShowbackKind(showbackKind string) *taikunshowback.EShowbackType {
	model, _ := ShowbackKinds.Get(showbackKind).(taikunshowback.EShowbackType)
	return &model
}

//import (
//	"github.com/itera-io/taikun-cli/utils/gmap"
//	"github.com/itera-io/taikungoclient/models"
//)
//
//var ShowbackKinds = gmap.New(
//	map[string]interface{}{
//		"general":  models.EShowbackType(100),
//		"external": models.EShowbackType(200),
//	},
//)
//
//func GetShowbackKind(showbackKind string) models.EShowbackType {
//	model, _ := ShowbackKinds.Get(showbackKind).(models.EShowbackType)
//	return model
//}
