package types

import (
	"github.com/itera-io/taikun-cli/utils/gmap"
	taikuncore "github.com/itera-io/taikungoclient/client"
)

//import (
//	"github.com/itera-io/taikun-cli/utils/gmap"
//	"github.com/itera-io/taikungoclient/models"
//)
//
//var SecurityGroupProtocols = gmap.New(
//	map[string]interface{}{
//		"icmp": models.SecurityGroupProtocol(100),
//		"tcp":  models.SecurityGroupProtocol(200),
//		"udp":  models.SecurityGroupProtocol(300),
//	},
//)
//
//func GetSecurityGroupProtocol(protocol string) models.SecurityGroupProtocol {
//	model, _ := SecurityGroupProtocols.Get(protocol).(models.SecurityGroupProtocol)
//	return model
//}

var SecurityGroupProtocols = gmap.New(
	map[string]interface{}{
		"icmp": taikuncore.SECURITYGROUPPROTOCOL_ICMP,
		"tcp":  taikuncore.SECURITYGROUPPROTOCOL_TCP,
		"udp":  taikuncore.SECURITYGROUPPROTOCOL_UDP,
	},
)

func GetSecurityGroupProtocol(protocol string) taikuncore.SecurityGroupProtocol {
	model, _ := SecurityGroupProtocols.Get(protocol).(taikuncore.SecurityGroupProtocol)
	return model
}
