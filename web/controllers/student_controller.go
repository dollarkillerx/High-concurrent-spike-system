/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-17
* Time: 上午10:54
* */
package controllers

import (
	"High-concurrent-spike-system/repositories"
	"High-concurrent-spike-system/services"
	"github.com/kataras/iris/mvc"
)

type StudentController struct {

}

func (s *StudentController) GetData() mvc.View {
	dao := repositories.NewStudentManger()
	mange := services.NewStudentServiceMange(dao)
	data := mange.ShowData()
	return mvc.View{
		Name:"student/index.html",
		Data:data,
	}
}
