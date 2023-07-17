package concurrent

import (
	"reflect"
	"testing"
)

func Test_ConcurrentDictionary(t *testing.T) {
	dict := NewDictionary[int, string]()

	count := 5000
	for i := 0; i < count; i++ {
		go func() {
			dict.Set(1, "asdf")
			print(dict.Get(1))
		}()
	}
}

func TestReplace(t *testing.T) {
	dict := NewDictionary[string, int]()

	dict.Set("apple", 5)
	dict.Set("banana", 10)

	newData := map[string]int{"orange": 7, "grape": 3}

	dict.Replace(newData)

	expectedData := map[string]int{"orange": 7, "grape": 3}
	if !reflect.DeepEqual(dict.data, expectedData) {
		t.Errorf("Replace method failed: expected %v, got %v", expectedData, dict.data)
	}
}
