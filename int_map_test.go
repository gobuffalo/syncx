package syncx

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_IntMap(t *testing.T) {
	r := require.New(t)

	sm := &IntMap{}

	sm.Store("a", 0)

	s, ok := sm.Load("a")
	r.True(ok)
	r.Equal(0, s)

	s, ok = sm.LoadOrStore("b", 1)
	r.True(ok)
	r.Equal(1, s)

	s, ok = sm.LoadOrStore("b", -1)
	r.True(ok)
	r.Equal(1, s)

	var keys []string

	sm.Range(func(key string, value int) bool {
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

	func(m *IntMap) {
		m.Store("c", 2)
	}(sm)
	s, ok = sm.Load("c")
	r.True(ok)
	r.Equal(2, s)
}
