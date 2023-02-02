package concurrent

import "sync"

// Dictionary The keys of the map must be of a comparable type such as bool, string, numeric types (int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128), array (when elements are comparable), or struct (when all fields are comparable).
type Dictionary[TK comparable, T any] struct {
	sync.RWMutex
	data map[TK]T
}

func NewDictionary[TK comparable, T any]() *Dictionary[TK, T] {
	dictionary := Dictionary[TK, T]{}
	dictionary.Init()
	return &dictionary
}

func (d *Dictionary[TK, T]) Init() {
	d.data = make(map[TK]T)
}

func (d *Dictionary[TK, T]) Get(key TK) (T, bool) {
	d.RLock()
	defer d.RUnlock()
	value, ok := d.data[key]
	return value, ok
}

func (d *Dictionary[TK, T]) Set(key TK, value T) {
	d.Lock()
	defer d.Unlock()
	d.data[key] = value
}
