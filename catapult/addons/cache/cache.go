package cache

import "github.com/BorisBorshevsky/GolangDemos/catapult"

func Add(provider Provider) *cacheMaker {
	return &cacheMaker{
		Provider: provider,
	}
}

func (c *cacheMaker) WithKey(cacheKey string) *cacheMaker {
	c.key = cacheKey
	return c
}

func (c *cacheMaker) extractKey(request *catapult.Request) string {
	if c.key != "" {
		return c.key
	}

	return request.Raw().URL.String()
}

func (c *cacheMaker) Register(ctx *catapult.Ctx) {
	ctx.AddJustBefore(func(request *catapult.Request) {
		key := keyPrefix + c.extractKey(request)
		ctx.Data[ctxKey] = key
		res, err := c.Get(key)

		if err == nil {
			ctx.Data[ctxValue] = res
			ctx.SkipDispatch()
		}

	})

	ctx.AddAfter(func(response *catapult.Response) {
		if ctx.Data[ctxValue] != nil {
			response.Write(ctx.Data[ctxValue].([]byte))
		} else {
			if response.Valid() && !response.Context.WasDispatchSkipped() {
				key := ctx.Data[ctxKey].(string)
				c.Set(key, string(response.Bytes()))
			}
		}
	})
}
