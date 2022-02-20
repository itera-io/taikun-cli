package types

import (
	"github.com/itera-io/taikun-cli/utils/gmap"
	"github.com/itera-io/taikungoclient/models"
)

var SecurityGroupProtocols = gmap.New(
	map[string]interface{}{
		"icmp": models.SecurityGroupProtocol(100),
		"tcp":  models.SecurityGroupProtocol(200),
		"udp":  models.SecurityGroupProtocol(300),
	},
)

func GetSecurityGroupProtocol(protocol string) models.SecurityGroupProtocol {
	model, _ := SecurityGroupProtocols.Get(protocol).(models.SecurityGroupProtocol)
	return model
}
