package types

import "strconv"

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
