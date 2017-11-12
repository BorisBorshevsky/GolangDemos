package cache

import (
	"github.com/BorisBorshevsky/GolangDemos/catapult"
	"github.com/k0kubun/pp"
)

const fallbackKeyPrefix = "$fb$"

type falbackCacheMaker struct {
	Provider
	key string
}

func AddFallbackCache(provider Provider) *falbackCacheMaker {
	return &falbackCacheMaker{
		Provider: provider,
	}
}

func (c *falbackCacheMaker) WithKey(cacheKey string) *falbackCacheMaker {
	c.key = cacheKey
	return c
}

func (c *falbackCacheMaker) extractKey(request *catapult.Request) string {
	if c.key != "" {
		return c.key
	}

	return request.Raw().URL.String()
}

func (c *falbackCacheMaker) Register(ctx *catapult.Ctx) {
	ctx.AddJustBefore(func(request *catapult.Request) {
		key := fallbackKeyPrefix + c.extractKey(request)
		ctx.Data[fbCtxKey] = key
	})

	ctx.AddAfter(func(response *catapult.Response) {
		key := ctx.Data[fbCtxKey].(string)
		if response.Valid() {
			if !response.Context.WasDispatchSkipped() {
				c.Set(key, string(response.Bytes()))
			}
		} else {
			pp.Println("HERE")
			cacheResp, err := c.Get(key)
			if err != nil {
				pp.Println("FB CAHCE??")
				if err != NotExist {
					pp.Println("FB CACHE ERROR", err)
					return
				}
			} else {
				pp.Println("WE HERE")
				response.Write(cacheResp)
				response.Context.SetError(nil)
			}
		}
	})
}
