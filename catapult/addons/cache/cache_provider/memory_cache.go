package cacheProvider

import "github.com/BorisBorshevsky/GolangDemos/catapult/addons/cache"

var memoryCache = &MemoryCacheProvider{}

func MemoryCache() *MemoryCacheProvider {
	return memoryCache
}

type MemoryCacheProvider map[string][]byte

func (*MemoryCacheProvider) Set(k, v string) error {
	(*memoryCache)[k] = []byte(v)
	return nil
}

func (*MemoryCacheProvider) Get(k string) ([]byte, error) {
	val, ex := (*memoryCache)[k]
	if !ex {
		return []byte{}, cache.NotExist
	}
	return val, nil
}
