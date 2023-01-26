package boolean_helper

import "testing"

func TestBoolToIntTrue(t *testing.T) {
	expected := 1
	actual := ToInt(true)
	if expected != actual {
		panic("wrong conversion")
	}
}

func TestBoolToIntFalse(t *testing.T) {
	expected := 0
	actual := ToInt(false)
	if expected != actual {
		panic("wrong conversion")
	}
}
