/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-17
* Time: 上午11:17
* */
package main

import (
	"High-concurrent-spike-system/config"
	"github.com/kataras/iris"
	"github.com/kataras/iris/view"
	"log"
)

func main() {
	// 1.创建iris 实例
	app := iris.New()

	// 2.设置错误等级
	if config.Config.Debug {
		app.Logger().SetLevel("debug")
	}

	// 3.注册模板
	var template *view.HTMLEngine
	if config.Config.Debug {
		template = iris.HTML("./web/views",".html").Layout("shared/layout.html").Reload(true) // .Reload 热加载
	}else{
		template = iris.HTML("./web/views",".html").Layout("shared/layout.html").Reload(false) // .Reload 热加载
	}
	app.RegisterView(template)

	// 4.设置静态文件目录
	app.StaticWeb("/assets","./web/assets") // 参数1:url前缀,参数二本地地址

	// 5.设置异常页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message",ctx.Values().GetStringDefault("message","访问页面出错"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	// 6.注册控制器


	// 7.启动服务
	err := app.Run(
		iris.Addr(config.Config.Host), // 地址
		iris.WithCharset("UTF-8"), // 国际化
		iris.WithOptimizations, // 自动优化
		iris.WithoutServerError(iris.ErrServerClosed), // 忽略框架错误
	)

	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}

}

