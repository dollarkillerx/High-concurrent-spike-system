package middleware

import (
	"High-concurrent-spike-system/config"
	"github.com/kataras/iris"
)

var (
	ints chan int
)

func init() {
	ints = make(chan int, config.MyConfig.App.MaxRequest)
}

func GlobalAfter(ctx iris.Context) {
	if len(ints) < config.MyConfig.App.MaxRequest {
		ints <- 1
		ctx.Next()
	} else {
		ctx.StatusCode(503)
		ctx.JSON(iris.Map{"Code": 503, "Msg": "服务器达到最大速率"})
		return
	}
	<-ints
}
