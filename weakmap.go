// Package weakmap implements a map with weak keys.
//
// One use for such a map is storing auxiliary data with values that are not
// under your control. For example, if you wanted to associated additional
// information with values of some type *p.T, where package p is not under your
// control, you could use a weakmap whose keys are *p.T values and whose values
// are your additional information. // The advantage over using an ordinary map
// is that when a *p.T gets garbage-collected, your associated data is removed
// from the map.
//
// This should be considered a proof of concept, for educational purposes only.
// The problem is that the implementation requires setting a finalizer on the
// keys, which removes any existing finalizer. There is no way for the program
// to tell if there is an existing finalizer, so it cannot detect this problem.
package weakmap

import (
	"reflect"
	"runtime"
	"sync"
)

// A Map is a map with weak keys.
// A zero Map is ready for use.
type Map struct {
	mu sync.Mutex
	m  map[uintptr]interface{}
}

// Put associates key with value in the map. The key must be a pointer.
func (m *Map) Put(key, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.m == nil {
		m.m = map[uintptr]interface{}{}
	}
	runtime.SetFinalizer(key, m.removeKey)
	// Convert the key to a uintptr to hide it from the garbage collector.
	m.m[toInt(key)] = value
}

func (m *Map) Get(key interface{}) interface{} {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.m[toInt(key)]
}

func (m *Map) Len() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.m)
}

func (m *Map) removeKey(key interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.m, toInt(key))
}

func toInt(key interface{}) uintptr {
	return reflect.ValueOf(key).Pointer()
}
