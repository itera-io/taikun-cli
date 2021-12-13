package cmdutils

import (
	"strings"

	"github.com/itera-io/taikungoclient/models"
)

func GetUserRole(role string) models.UserRole {
	role = strings.ToLower(role)
	if role == "user" {
		return 400
	}
	// manager
	return 200
}
