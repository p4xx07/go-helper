package string_helper

import "strings"

func Truncate(text string, length int) string {
	const dots string = "..."
	if len(text) < length {
		return strings.TrimSpace(text)
	}
	return strings.TrimSpace(text[:length]) + dots
}
