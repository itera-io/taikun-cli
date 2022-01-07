package types

import (
	"strings"

	"github.com/itera-io/taikungoclient/models"
)

var ServerRoles = map[string]interface{}{
	"bastion":    models.CloudRole(100),
	"kubemaster": models.CloudRole(200),
	"kubeworker": models.CloudRole(300),
}

func GetServerRole(serverRole string) models.CloudRole {
	return ServerRoles[strings.ToLower(serverRole)].(models.CloudRole)
}
