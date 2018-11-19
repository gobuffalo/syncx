package syncx

import (
	"sync"
)

// asdfasdf
type IntMap struct {
	data sync.Map
}

func (m *IntMap) Delete(key string) {
	m.data.Delete(key)
}

func (m *IntMap) Load(key string) (int, bool) {
	i, ok := m.data.Load(key)
	if !ok {
		return 0, false
	}
	s, ok := i.(int)
	return s, ok
}

func (m *IntMap) LoadOrStore(key string, value int) (int, bool) {
	i, ok := m.data.LoadOrStore(key, value)
	s, ok := i.(int)
	return s, ok
}

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

func (m *IntMap) Store(key string, value int) {
	m.data.Store(key, value)
}
