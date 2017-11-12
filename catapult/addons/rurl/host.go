package rurl

import "github.com/BorisBorshevsky/GolangDemos/catapult"

type hostAddon string

func (p hostAddon) Register(ctx *catapult.Ctx) {
	ctx.AddBefore(func(request *catapult.Request) {
		request.Raw().URL.Host = string(p)
	})
}

func Host(host string) hostAddon {
	return hostAddon(host)
}
