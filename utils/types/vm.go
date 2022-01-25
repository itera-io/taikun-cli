package types

func GetVMRebootType(hardReboot bool) string {
	if hardReboot {
		return "HARD"
	}
	return "SOFT"
}
