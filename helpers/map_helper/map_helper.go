package map_helper

func Keys[TK comparable, T any](m map[TK]T) []TK {
	output := make([]TK, 0, len(m))
	for key := range m {
		output = append(output, key)
	}
	return output
}

func Values[TK comparable, T any](m map[TK]T) []T {
	output := make([]T, 0, len(m))
	for _, value := range m {
		output = append(output, value)
	}
	return output
}
