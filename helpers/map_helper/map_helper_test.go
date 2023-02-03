package map_helper

import (
	"reflect"
	"sort"
	"testing"
)

func Test_MapKeys(t *testing.T) {
	input := map[int]string{
		1: "ciao",
		2: "ciao",
		3: "test",
	}

	expectedOutput := []int{1, 2, 3}
	output := Keys[int, string](input)
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Keys(%v) = %v, expected %v", input, output, expectedOutput)
	}
}

func Test_MapValues(t *testing.T) {
	input := map[int]string{
		1: "ciao",
		2: "ciao",
		3: "test",
	}

	expectedOutput := []string{"ciao", "ciao", "test"}
	output := Values[int, string](input)
	sort.Strings(output)
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Keys(%v) = %v, expected %v", input, output, expectedOutput)
	}
}
