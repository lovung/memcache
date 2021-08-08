package memcache

import (
	"testing"
	"time"
)

func TestFull(t *testing.T) {
	mc := NewCache()
	mc.SetUntil("key", "value", 0)
	mc.Set("key", "value")
	v := mc.Get("key")
	if "value" != v.(string) {
		t.Errorf("Get failed")
	}
	mc.Del("key")
	v = mc.Get("key")
	if v != nil {
		t.Errorf("Get failed")
	}
	mc.Set("key", "value")
	v = mc.Take("key")
	if "value" != v.(string) {
		t.Errorf("Get failed")
	}
	v = mc.Get("key")
	if v != nil {
		t.Errorf("Get failed")
	}

	mc.SetUntil("key", "value", 1*time.Millisecond)
	time.Sleep(2 * time.Millisecond)
	v = mc.Get("key")
	if v != nil {
		t.Errorf("Get failed")
	}
	v = mc.Take("key")
	if v != nil {
		t.Errorf("Get failed")
	}

	mc.SetUntil("key", "value", 1*time.Millisecond)
	time.Sleep(2 * time.Millisecond)
	v = mc.Take("key")
	if v != nil {
		t.Errorf("Get failed")
	}
}

func BenchmarkSet(b *testing.B) {
	mc := NewCache()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mc.Set("key", "value")
	}
}

func BenchmarkGet(b *testing.B) {
	mc := NewCache()
	mc.Set("key", "value")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mc.Get("key")
	}
}

func BenchmarkSetTake(b *testing.B) {
	mc := NewCache()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mc.Set("key", "value")
		mc.Take("key")
	}
}

func BenchmarkSetGetDel(b *testing.B) {
	mc := NewCache()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mc.Set("key", "value")
		mc.Get("key")
		mc.Del("key")
	}
}
