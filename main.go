package main

import (
	"LotteryProject/bootstrap"
	"LotteryProject/web/middleware/identity"
	"LotteryProject/web/routes"
	"fmt"
)

var port = 8080

func NewApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("GO抽奖系统", "YuSheng")

	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)

	return app
}

func main() {
	app := NewApp()
	app.Listen(fmt.Sprintf(":%d", port))
}
