package utils

import (
	"strconv"
	"strings"

	"github.com/itera-io/taikungoclient/models"
)

const UnlockedMode = "unlock"
const LockedMode = "lock"

func GetUserRole(role string) models.UserRole {
	role = strings.ToLower(role)
	if role == "user" {
		return 400
	}
	if role == "manager" {
		return 200
	}
	return -1
}

func Atoi32(str string) (int32, error) {
	res, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(res), nil
}

func GiBToMiB(gibiBytes float64) int32 {
	return int32(gibiBytes * 1024)
}
