package rurl

import "github.com/BorisBorshevsky/GolangDemos/catapult"

type pathAddon string

func (p pathAddon) Register(ctx *catapult.Ctx) {
	ctx.AddBefore(func(request *catapult.Request) {
		request.Raw().URL.Path = string(p)
	})
}

func Path(path string) pathAddon {
	return pathAddon(path)
}
