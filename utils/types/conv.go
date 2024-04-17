package types

import "strconv"

func Atoi32(str string) (int32, error) {
	res, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(res), nil
}

func GiBToB(gibiBytes int32) float64 {
	return float64(1073741824) * float64(gibiBytes)
}
