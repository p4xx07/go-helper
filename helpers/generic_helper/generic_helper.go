package generic_helper

import (
	"errors"
	"reflect"
	"strconv"
	"time"
)

func ConvertFromString[T any](in string) (*T, error) {
	var out T
	r := reflect.ValueOf(&out)
	e := r.Elem()

	switch e.Kind() {
	case reflect.String:
		e.SetString(in)
		return &out, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		o, err := strconv.ParseInt(in, 10, 64)
		if err != nil {
			return nil, err
		}
		e.SetInt(o)
		return &out, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u, err := strconv.ParseUint(in, 10, 64)
		if err != nil {
			return nil, err
		}
		e.SetUint(u)
		return &out, nil
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(in, 64)
		if err != nil {
			return nil, err
		}
		e.SetFloat(f)
		return &out, nil
	case reflect.Struct:
		if e.Type() == reflect.TypeOf(time.Time{}) {
			t, err := time.Parse(time.RFC3339, in)
			if err != nil {
				return nil, err
			}
			e.Set(reflect.ValueOf(t))
			return &out, nil
		} else {
			return nil, errors.New("not supported struct type")
		}
	default:
		return nil, errors.New("could not convert in a valid type")
	}
}
