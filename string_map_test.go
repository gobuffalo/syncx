package syncx

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_StringMap(t *testing.T) {
	r := require.New(t)

	sm := &StringMap{}

	sm.Store("a", "A")

	s, ok := sm.Load("a")
	r.True(ok)
	r.Equal("A", s)

	s, ok = sm.LoadOrStore("b", "B")
	r.True(ok)
	r.Equal("B", s)

	s, ok = sm.LoadOrStore("b", "BB")
	r.True(ok)
	r.Equal("B", s)

	var keys []string

	sm.Range(func(key string, value string) bool {
		keys = append(keys, key)
		return true
	})

	sort.Slice(keys, func(a, b int) bool {
		return keys[a] < keys[b]
	})

	r.Equal([]string{"a", "b"}, keys)

	sm.Delete("b")
	_, ok = sm.Load("b")
	r.False(ok)

	func(m *StringMap) {
		m.Store("c", "C")
	}(sm)
	s, ok = sm.Load("c")
	r.True(ok)
	r.Equal("C", s)
}
