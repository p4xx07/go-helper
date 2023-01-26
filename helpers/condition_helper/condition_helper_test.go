package condition_helper

import (
	"testing"
)

func Test_TernaryTrue(t *testing.T) {
	expected := 200
	actual := Ternary(1 == 1, 200, 500)
	if expected != actual {
		panic("ternary not working")
	}
}

func Test_TernaryFalse(t *testing.T) {
	expected := 500
	actual := Ternary(1 != 1, 200, 500)
	if expected != actual {
		panic("ternary not working")
	}
}

type A struct {
	B string
}

func Test_TernaryTrueWithStruct(t *testing.T) {
	expected := "success"
	actual := Ternary(1 == 1, A{B: "success"}, A{B: "fail"})
	if expected != actual.B {
		panic("ternary not working")
	}
}

func Test_TernaryPredicate(t *testing.T) {
	expected := 200
	actual := TernaryPredicate(func() bool {
		operation := 1 + 1
		return operation == 2
	}, 200, 500)

	if expected != actual {
		panic("ternary not working")
	}
}
