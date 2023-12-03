package memcache

import "time"

type shard[K Stringer, V any] struct {
	hashmap map[K]V
	// cache   *cache // use if neccessary
}

func (s *shard[K Stringer, V any]) set(k K, v V, d time.Duration) {
	s.hashmap[k] = item{
		value:     v,
		expiredAt: time.Now().Add(d),
	}
}

func (s *shard[K Stringer, V any]) get(k K) V {
	item, ok := s.hashmap[k]
	if !ok {
		return nil
	}
	if item.expiredAt.Before(time.Now()) {
		delete(s.hashmap, k)
		return nil
	}
	return item.value
}

func (s *shard[K Stringer, V any]) take(k K) V {
	item, ok := s.hashmap[k]
	if !ok {
		return nil
	}
	delete(s.hashmap, k)
	if item.expiredAt.Before(time.Now()) {
		return nil
	}
	return item.value
}

func (s *shard[K Stringer, V any]) delete(k K) {
	delete(s.hashmap, k)
}
