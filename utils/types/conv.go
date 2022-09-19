package types

import "strconv"

func Atoi32(str string) (int32, error) {
	res, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(res), nil
}

func GiBToMiB(gibiBytes float64) float64 {
	return gibiBytes * 1024
}

func GiBToB(gibiBytes int) int64 {
	return int64(int64(gibiBytes) * 1024 * 1024 * 1024)
}
