# memcache
Memory Caching in Golang. 
Preparing for Generics in go1.18.

## Features

- [x] Set
- [x] SetUntil
- [x] Get
- [x] Take
- [x] Delete
- [ ] GC

## Config

- [ ] Size
- [ ] Shard
- [ ] TTL


## Benchmark
## Env
```
goos: darwin
goarch: amd64
pkg: github.com/lovung/memcache
cpu: Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz
```

## Results
```
BenchmarkSet-8   	     7606341	       132.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet-8   	     11004312	       114.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkSetTake-8   	 4407631	       267.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkSetGetDel-8   	 4247701	       276.4 ns/op	       0 B/op	       0 allocs/op
```
