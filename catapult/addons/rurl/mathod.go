package rurl

import "github.com/BorisBorshevsky/GolangDemos/catapult"

type methodAddon string

func (p methodAddon) Register(ctx *catapult.Ctx) {
	ctx.AddBefore(func(request *catapult.Request) {
		request.Raw().Method = string(p)
	})
}

func Method(method string) methodAddon {
	return methodAddon(method)
}
