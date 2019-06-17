/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-16
* Time: 下午11:14
* */
package controllers

import (
	"High-concurrent-spike-system/repositories"
	"High-concurrent-spike-system/services"
	"github.com/kataras/iris/mvc"
)

type MovieController struct {

}

func (c *MovieController) Get() mvc.View {
	// 创建数据库操作对象
	movieRepository := repositories.NewMovieManager()
	// 建立service
	movieService := services.NewMovieServiceManger(movieRepository)
	name := movieService.ShowMovieName()

	return mvc.View{
		Name:"movie/index.html",
		Data:name,
	}
}


