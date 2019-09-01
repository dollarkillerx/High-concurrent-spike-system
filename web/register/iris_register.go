package register

import (
	"High-concurrent-spike-system/config"
	"High-concurrent-spike-system/web/middleware"
	"High-concurrent-spike-system/web/router"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	recover2 "github.com/kataras/iris/middleware/recover"
)

func Iris() *iris.Application {
	app := iris.New()

	app.Use(recover2.New())

	if config.MyConfig.App.Debug {
		app.Use(logger.New())
	}

	// 设置错误等级
	app.Logger().SetLevel(config.MyConfig.App.LogLevel)

	// 注册视图
	app.RegisterView(iris.HTML("./views", ".html"))

	// 注册全局流量控制
	app.Use(middleware.GlobalAfter)

	// 注册路由
	router.Register(app)

	return app
}
