package syncx

import (
	"sync"
)

// IntMap wraps sync.Map and uses the following types:
// key:   string
// value: int
type IntMap struct {
	data sync.Map
}

// Delete the key from the map
func (m *IntMap) Delete(key string) {
	m.data.Delete(key)
}

// Load the key from the map.
// Returns int or bool.
// A false return indicates either the key was not found
// or the value is not of type int
func (m *IntMap) Load(key string) (int, bool) {
	i, ok := m.data.Load(key)
	if !ok {
		return 0, false
	}
	s, ok := i.(int)
	return s, ok
}

// LoadOrStore will return an existing key or
// store the value if not already in the map
func (m *IntMap) LoadOrStore(key string, value int) (int, bool) {
	i, _ := m.data.LoadOrStore(key, value)
	s, ok := i.(int)
	return s, ok
}

// Range over the int values in the map
func (m *IntMap) Range(f func(key string, value int) bool) {
	m.data.Range(func(k, v interface{}) bool {
		key, ok := k.(string)
		if !ok {
			return false
		}
		value, ok := v.(int)
		if !ok {
			return false
		}
		return f(key, value)
	})
}

// Store a int in the map
func (m *IntMap) Store(key string, value int) {
	m.data.Store(key, value)
}
