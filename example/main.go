package main

import (
	"fmt"
	"time"

	"github.com/lovung/memcache"
)

func main() {
	mc := memcache.NewCache()
	mc.SetUntil("key", "value", 0)
	mc.Set("key", "value")
	v := mc.Get("key")
	if "value" != v.(string) {
		fmt.Println("Get failed")
	}
	mc.Del("key")
	v = mc.Get("key")
	if v != nil {
		fmt.Println("Get failed")
	}
	mc.Set("key", "value")
	v = mc.Take("key")
	if "value" != v.(string) {
		fmt.Println("Get failed")
	}
	v = mc.Get("key")
	if v != nil {
		fmt.Println("Get failed")
	}

	mc.SetUntil("key", "value", 1*time.Millisecond)
	time.Sleep(2 * time.Millisecond)
	v = mc.Get("key")
	if v != nil {
		fmt.Println("Get failed")
	}
	v = mc.Take("key")
	if v != nil {
		fmt.Println("Get failed")
	}

	mc.SetUntil("key", "value", 1*time.Millisecond)
	time.Sleep(2 * time.Millisecond)
	v = mc.Take("key")
	if v != nil {
		fmt.Println("Get failed")
	}

	fmt.Println("Success")
}
