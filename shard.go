package memcache

import "time"

type shard struct {
	hashmap map[string]item
	// cache   *cache // use if neccessary
}

func (s *shard) set(k string, v interface{}, d time.Duration) {
	s.hashmap[k] = item{
		value:     v,
		expiredAt: time.Now().Add(d),
	}
}

func (s *shard) get(k string) interface{} {
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

func (s *shard) take(k string) interface{} {
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

func (s *shard) delete(k string) {
	delete(s.hashmap, k)
}
