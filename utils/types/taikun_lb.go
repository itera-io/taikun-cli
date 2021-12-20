package types

func IsInRouterIDRange(x int32) bool {
	return x >= 1 && x <= 255
}
