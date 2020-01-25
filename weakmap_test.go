package weakmap

import (
	"runtime"
	"testing"
	"time"
)

func Test(t *testing.T) {
	var m Map
	key1 := newkey(1)
	key2 := newkey(2)
	got := m.Get(key1)
	if got != nil {
		t.Error("want nil")
	}

	m.Put(key1, "one")
	m.Put(key2, "two")
	if got := m.Get(key1); got != "one" {
		t.Error("fail one")
	}
	if got := m.Get(key2); got != "two" {
		t.Error("fail two")
	}

	key1 = nil
	runtime.GC()
	time.Sleep(2 * time.Second)
	if got := m.Len(); got != 1 {
		t.Error("fail want 1")
	}
	if got := m.Get(key2); got != "two" {
		t.Error("fail two")
	}
}

func newkey(i int) *int { return &i }
