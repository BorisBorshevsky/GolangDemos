package rurl

import "github.com/BorisBorshevsky/GolangDemos/catapult"

type params struct {
	keys map[string]string
}

func AddParam(key, val string) *params {
	return &params{
		keys: map[string]string{key: val},
	}
}

func (p *params) Register(ctx *catapult.Ctx) {
	ctx.AddBefore(func(request *catapult.Request) {
		params := request.Raw().URL.Query()
		for k, v := range p.keys {
			params.Add(k, v)
		}
		request.Raw().URL.RawQuery = params.Encode()
	})
}
