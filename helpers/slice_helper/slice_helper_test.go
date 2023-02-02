package slice_helper

import (
	"reflect"
	"testing"
)

func TestRemoveEmptyEntries(t *testing.T) {
	input := []string{"", "hello", "", "world", ""}
	expectedOutput := []string{"hello", "world"}
	output := RemoveEmptyEntries(input)
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("RemoveEmptyEntries(%v) = %v, expected %v", input, output, expectedOutput)
	}
}
func TestMapInt(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expectedOutput := []int{2, 4, 6, 8, 10}
	output := Map(&input, func(l int) int {
		return l * 2
	})
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Map(%v) = %v, expected %v", input, output, expectedOutput)
	}
}

func TestMapChangeType(t *testing.T) {
	input := []string{"foo", "ciao", "world"}
	expectedOutput := []int{3, 4, 5}
	output := Map[string, int](&input, func(l string) int {
		return len(l)
	})
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Map(%v) = %v, expected %v", input, output, expectedOutput)
	}
}

func TestDistinct(t *testing.T) {
	input := []string{"foo", "ciao", "world", "world", "world", "lol", "ciao"}
	expectedOutput := []string{"foo", "ciao", "world", "lol"}
	output := Distinct[string](input)
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Distinct(%v) = %v, expected %v", input, output, expectedOutput)
	}
}
