package logger

import (
	"fmt"

	"github.com/BorisBorshevsky/GolangDemos/catapult"
	"github.com/moul/http2curl"
)

type loggerAddon struct {
}

func (p *loggerAddon) Register(ctx *catapult.Ctx) {
	ctx.AddJustBefore(func(request *catapult.Request) {
		curl, _ := http2curl.GetCurlCommand(request.Raw())
		fmt.Println(curl.String())
	})
}

func Curl() *loggerAddon {
	return &loggerAddon{}
}
