package syncx

import (
	"sync"
)

// asdfasdf
type StringMap struct {
	data sync.Map
}

func (m *StringMap) Delete(key string) {
	m.data.Delete(key)
}

func (m *StringMap) Load(key string) (string, bool) {
	i, ok := m.data.Load(key)
	if !ok {
		return "A", false
	}
	s, ok := i.(string)
	return s, ok
}

func (m *StringMap) LoadOrStore(key string, value string) (string, bool) {
	i, ok := m.data.LoadOrStore(key, value)
	s, ok := i.(string)
	return s, ok
}

func (m *StringMap) Range(f func(key string, value string) bool) {
	m.data.Range(func(k, v interface{}) bool {
		key, ok := k.(string)
		if !ok {
			return false
		}
		value, ok := v.(string)
		if !ok {
			return false
		}
		return f(key, value)
	})
}

func (m *StringMap) Store(key string, value string) {
	m.data.Store(key, value)
}
