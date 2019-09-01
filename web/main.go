package main

import (
	"High-concurrent-spike-system/config"
	"High-concurrent-spike-system/web/register"
	"github.com/kataras/iris"
)

func main() {
	app := register.Iris()

	app.Run(iris.Addr(config.MyConfig.App.Host),iris.WithCharset("UTF-8"))
}
