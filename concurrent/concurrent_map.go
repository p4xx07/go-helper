package concurrent

import "sync"

// Dictionary The keys of the map must be of a comparable type such as bool, string, numeric types (int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128), array (when elements are comparable), or struct (when all fields are comparable).
type Dictionary[TK comparable, T any] struct {
	mutex sync.RWMutex
	data  map[TK]T
}

func NewDictionary[TK comparable, T any]() *Dictionary[TK, T] {
	dictionary := Dictionary[TK, T]{}
	dictionary.data = make(map[TK]T)
	return &dictionary
}

func (d *Dictionary[TK, T]) Get(key TK) (T, bool) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	value, ok := d.data[key]
	return value, ok
}

func (d *Dictionary[TK, T]) Set(key TK, value T) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.data[key] = value
}

func (d *Dictionary[TK, T]) Length() int {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return len(d.data)
}

func (d *Dictionary[TK, T]) ContainsKey(key TK) bool {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	_, ok := d.data[key]
	return ok
}

func (d *Dictionary[TK, T]) Values() []T {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	values := make([]T, 0, len(d.data))
	for _, value := range d.data {
		values = append(values, value)
	}
	return values
}

func (d *Dictionary[TK, T]) Copy() *Dictionary[TK, T] {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	newDict := NewDictionary[TK, T]()
	for key, value := range d.data {
		newDict.Set(key, value)
	}
	return newDict
}

func (d *Dictionary[TK, T]) Keys() []TK {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	keys := make([]TK, 0, len(d.data))
	for key := range d.data {
		keys = append(keys, key)
	}
	return keys
}

func (d *Dictionary[TK, T]) Delete(key TK) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	delete(d.data, key)
}

func (d *Dictionary[TK, T]) Clear() {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.data = make(map[TK]T)
}

func (d *Dictionary[TK, T]) IsEmpty() bool {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return len(d.data) == 0
}

func (d *Dictionary[TK, T]) Replace(newData map[TK]T) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.data = newData
}
