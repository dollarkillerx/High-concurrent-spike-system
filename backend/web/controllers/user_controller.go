/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-17
* Time: 下午9:51
* */
package controllers

import (
	"High-concurrent-spike-system/repositories"
	"database/sql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type UserController struct {

}

func (u *UserController) GetBy(ctx iris.Context,name string) mvc.Result {
	userRepositories := repositories.NewUserRepositories()
	user, e := userRepositories.GetByName(name)
	if e == sql.ErrNoRows {

	}else if e != nil {
		return
	}
	return mvc.View{
		Name:"user/user.html",
		Data:user,
	}
}
