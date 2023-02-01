package concurrent

import "testing"

func Test_ConcurrentDictionary(t *testing.T) {
	dict := Dictionary[int, string]{}
	dict.Init()

	count := 5000
	for i := 0; i < count; i++ {
		go func() {
			dict.Set(1, "asdf")
			print(dict.Get(1))
		}()
	}
}
