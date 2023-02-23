package type_helper

import "testing"

func TestSafeCastOk(t *testing.T) {
	var obj any = A{B: "test"}
	casted, ok := SafeCast[A](obj, nil)

	if !ok {
		panic("not ok")
	}

	if casted.B != "test" {
		panic("not the same")
	}
}

func TestSafeCastNotOk(t *testing.T) {
	var obj any = A{B: "test"}
	casted, ok := SafeCast[string](obj, nil)

	if ok {
		panic(" should not be ok")
	}

	if casted != nil {
		panic("should be nil")
	}
}

type A struct {
	B string
}
