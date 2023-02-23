package type_helper

func SafeCast[T any](obj any, defaultValue *T) (*T, bool) {
	result, ok := obj.(T)
	if !ok {
		return defaultValue, false
	}
	return &result, true
}
