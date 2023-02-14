package string_helper

import (
	"fmt"
	"testing"
)

func Test_Truncate(t *testing.T) {
	expected := "very lo..."
	actual := Truncate("very long text", 10)
	if expected != actual {
		fmt.Println("Actual: " + actual)
		panic("strings not equal")
	}
}
