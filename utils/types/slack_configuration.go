package types

import (
	"github.com/itera-io/taikun-cli/utils/gmap"
	taikuncore "github.com/itera-io/taikungoclient/client"
)

var SlackTypes = gmap.New(
	map[string]interface{}{
		"alert":   taikuncore.SLACKTYPE_ALERT,
		"general": taikuncore.SLACKTYPE_GENERAL,
	},
)

func GetSlackType(slackType string) *taikuncore.SlackType {
	model, _ := SlackTypes.Get(slackType).(taikuncore.SlackType)
	return &model
}

//import (
//	"github.com/itera-io/taikun-cli/utils/gmap"
//	"github.com/itera-io/taikungoclient/models"
//)
//
//var SlackTypes = gmap.New(
//	map[string]interface{}{
//		"alert":   models.SlackType(100),
//		"general": models.SlackType(200),
//	},
//)
//
//func GetSlackType(slackType string) models.SlackType {
//	model, _ := SlackTypes.Get(slackType).(models.SlackType)
//	return model
//}
