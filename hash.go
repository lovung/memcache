package memcache

import "github.com/zeebo/xxh3"

func hash(b []byte) uint64 {
	return xxh3.Hash(b)
}

func hashString(s string) uint64 {
	return xxh3.HashString(s)
}
