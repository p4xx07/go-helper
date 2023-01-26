package boolean_helper

func ToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
