package routes

import (
	"LotteryProject/bootstrap"
	"LotteryProject/services"
	"LotteryProject/web/controllers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Configure(b *bootstrap.Bootstrapper) {
	userService := services.NewUserService()
	giftService := services.NewGiftService()
	codeService := services.NewCodeService()
	resultService := services.NewResultService()
	userDayService := services.NewUserdayService()
	blackIpService := services.NewBlackipService()

	app := iris.New()
	index := mvc.New(app.Party("/"))
	index.Register(userService,
		giftService,
		codeService,
		resultService,
		userDayService,
		blackIpService)
	index.Handle(new(controllers.IndexController))

}
