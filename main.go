/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-16
* Time: 下午10:47
* */
package main

import (
	"High-concurrent-spike-system/web/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug") // 设置错误等级 开发模式

	// 注册html模板
	app.RegisterView(iris.HTML("./web/views",".html"))

	// 注册控制器
	mvc.New(app.Party("/movie")).Handle(new(controllers.MovieController))
	mvc.New(app.Party("/student")).Handle(new(controllers.StudentController))

	app.Run(iris.Addr(":8085"),iris.WithCharset("UTF-8"),iris.WithOptimizations)
}

