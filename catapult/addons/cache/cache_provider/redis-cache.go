package cacheProvider

import (
	"time"

	"github.com/BorisBorshevsky/GolangDemos/catapult/addons/cache"
	"gopkg.in/redis.v5"
)

func RedisTTLCache(ttl time.Duration) *RedisTTLCacheProvider {
	client := redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1" + ":6379",
		Password:     "", // no password set
		MaxRetries:   3,
		PoolSize:     10,
		PoolTimeout:  time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		DialTimeout:  2 * time.Second,
	})

	return &RedisTTLCacheProvider{
		Client:  client,
		timeout: ttl,
	}
}

type RedisTTLCacheProvider struct {
	*redis.Client
	timeout time.Duration
}

func (r *RedisTTLCacheProvider) Set(k, v string) error {
	return r.Client.Set(k, v, r.timeout).Err()
}

func (r *RedisTTLCacheProvider) Get(k string) ([]byte, error) {
	res, err := r.Client.Get(k).Bytes()
	if err == redis.Nil {
		return res, cache.NotExist
	}
	return res, err
}
