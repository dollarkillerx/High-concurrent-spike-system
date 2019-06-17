/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-17
* Time: 下午12:14
* */
package datamodels

type Product struct {
	Id          int64  `json:"id" sql:"ID" you:"id"`
	ProductName string `json:"product_name" sql:"productName" you:"ProductName"`
	ProductNum  int64  `json:"product_num" sql:"productNum" you:"ProductNum"`
	ProductImg  string `json:"product_img" sql:"productImg" you:"ProductImg"`
	ProductUrl  string `json:"product_url" sql:"productUrl" you:"ProductUrl"`
}
