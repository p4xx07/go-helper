package slice_helper

func Map[T any, TK any](input *[]T, f func(l T) TK) []TK {
	var output []TK
	for _, v := range *input {
		output = append(output, f(v))
	}
	return output
}

func Distinct[T comparable](input []T) []T {
	output := make([]T, 0, len(input))
	set := make(map[T]bool)
	for _, element := range input {
		if _, found := set[element]; !found {
			set[element] = true
			output = append(output, element)
		}
	}

	return output
}
