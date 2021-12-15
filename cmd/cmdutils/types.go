package cmdutils

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

func GetAWSRegion(region string) models.AwsRegion {
	switch region {
	case "us-east-1":
		return 1
	case "us-east-2":
		return 2
	case "us-west-1":
		return 3
	case "us-west-2":
		return 4
	case "eu-north-1":
		return 5
	case "eu-west-1":
		return 6
	case "eu-west-2":
		return 7
	case "eu-west-3":
		return 8
	case "eu-central-1":
		return 9
	case "eu-south-1":
		return 10
	case "ap-east-1":
		return 11
	case "ap-northeast-1":
		return 12
	case "ap-northeast-2":
		return 13
	case "ap-northeast-3":
		return 14
	case "ap-south-1":
		return 15
	case "ap-southeast-1":
		return 16
	case "ap-southeast-2":
		return 17
	case "sa-east-1":
		return 18
	case "us-gov-east-1":
		return 19
	case "us-gov-west-1":
		return 20
	case "cn-north-1":
		return 21
	case "cn-northwest-1":
		return 22
	case "ca-central-1":
		return 23
	case "me-south-1":
		return 24
	default: // af-south-1
		return 25
	}
}
