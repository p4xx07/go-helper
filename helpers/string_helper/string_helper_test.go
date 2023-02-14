package string_helper

import (
	"fmt"
	"testing"
)

func Test_Truncate(t *testing.T) {
	expected := "very long..."
	actual := Truncate("very long text", 10)
	if expected != actual {
		fmt.Println("Actual: " + actual)
		panic("strings not equal")
	}
}

func Test_TruncateOOB(t *testing.T) {
	Truncate("very", 10)
}
