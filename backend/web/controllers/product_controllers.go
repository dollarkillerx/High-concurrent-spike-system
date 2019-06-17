/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-17
* Time: 下午5:59
* */
package controllers

import (
	"High-concurrent-spike-system/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
)

type ProductController struct {
	Ctx iris.Context
	ProductService services.IProductService
}

func (p *ProductController) GetAll() mvc.View {
	products, e := p.ProductService.GetAllProduct()
	if e != nil {
		log.Println(e.Error())
		//p.Ctx.String()
	}
	return mvc.View{
		Name:"product/view.html",
		Data:iris.Map{
			"productArray":products,
		},
	}
}