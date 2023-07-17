package identity

import (
	"LotteryProject/bootstrap"
	"github.com/kataras/iris/v12"
	"time"
)

func New(b *bootstrap.Bootstrapper) iris.Handler {
	return func(ctx iris.Context) {
		// response headers
		ctx.Header("App-Name", b.AppName)
		ctx.Header("App-Owner", b.AppOwner)
		ctx.Header("App-Since", time.Since(b.AppSpawDate).String())

		ctx.Header("Server", "Iris: https://iris-go.com")

		ctx.ViewData("AppName", b.AppName)
		ctx.ViewData("AppOwner", b.AppOwner)
		ctx.Next()
	}
}

func Configure(b *bootstrap.Bootstrapper) {
	h := New(b)
	b.UseGlobal(h)
}
