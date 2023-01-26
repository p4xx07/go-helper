package slice_helper

func Map[T any, TK any](input *[]T, f func(l T) TK) []TK {
	var output []TK
	for _, v := range *input {
		output = append(output, f(v))
	}
	return output
}
