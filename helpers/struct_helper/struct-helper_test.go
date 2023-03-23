package struct_helper

import "testing"

type A struct {
	Number int    `json:"Number"`
	String string `json:"String"`
}

type B struct {
	Number int    `json:"Number"`
	String string `json:"String"`
}

func TestAsTypeJsonShouldNotFail(t *testing.T) {
	a := A{1, "test"}
	got, err := AsType[B](a)
	if err != nil {
		t.Errorf(err.Error())
	}
	want := B{1, "test"}
	if got.String != want.String || got.Number != want.Number {
		t.Errorf("convertion failed")
	}
}

func TestGetOrDefault(t *testing.T) {
	v := "hello"
	a := GetOrDefault[*string](nil, &v)

	if *a != "hello" {
		panic("failed")
	}

	v2 := "hello2"
	a = GetOrDefault[*string](&v, &v2)

	if *a != "hello" {
		panic("failed")
	}

	b := GetOrDefault("", "hello")
	if b != "hello" {
		panic("failed")
	}
}
