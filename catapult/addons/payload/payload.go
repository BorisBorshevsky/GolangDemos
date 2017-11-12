package payload

import "github.com/BorisBorshevsky/GolangDemos/catapult"

type payloadAddon struct {
	payload interface{}
}

func (p *payloadAddon) Register(ctx *catapult.Ctx) {
	ctx.AddBefore(func(request *catapult.Request) {
		request.Body(p.payload)
	})
}

func String(string string) *payloadAddon {
	return &payloadAddon{
		payload: []byte(string),
	}
}

func Json(json map[string]interface{}) *payloadAddon {
	return &payloadAddon{
		payload: json,
	}
}

func Add(payload interface{}) *payloadAddon {
	return &payloadAddon{
		payload: payload,
	}
}
