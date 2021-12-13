package cmdutils

import (
	"strconv"
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

func Atoi32(str string) (int32, error) {
	res, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(res), nil
}
