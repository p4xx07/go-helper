package struct_helper

import (
	"encoding/json"
)

func StructToString(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func AsType[T any](data any) (T, error) {
	var result T

	jsonData, err := json.Marshal(data)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetOrDefault[T comparable](value T, defaultValue T) T {
	if value != *new(T) {
		return value
	}
	return defaultValue
}
