package string_helper

func Truncate(text string, length int) string {
	const dots string = "..."
	return text[:length-len(dots)] + dots
}
