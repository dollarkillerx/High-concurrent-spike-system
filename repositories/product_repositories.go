/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-17
* Time: 下午12:34
* */
package repositories

import (
	"High-concurrent-spike-system/databases/mysql_conn"
	"High-concurrent-spike-system/datamodels"
	"errors"
)

// 第一步,向开发对应接口
// 第二步,实现定义接口

type IProduct interface {
	Insert(*datamodels.Product) (int64,error)
	Delete(int64) error
	Update(*datamodels.Product) error
	SelectByKey(int64) (*datamodels.Product,error)
	SelectAll() ([]*datamodels.Product,error)
}

type ProductManger struct {
	table string
}

func NewProductManger(table string) IProduct {
	return &ProductManger{
		table:table,
	}
}

func (p *ProductManger) Insert(product *datamodels.Product) (int64,error) {
	sql := "INSERT product SET product_name=?,product_num=?,product_image=?,product_url=?"
	stmt, e := mysql_conn.MySQL.Prepare(sql)
	if e != nil {
		return 0,e
	}
	result, e := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImg, product.ProductUrl)
	if e != nil {
		return 0,e
	}
	return result.LastInsertId()
}

func (p *ProductManger) Delete(id int64) error {
	sql := "DELETE FORM product WHERE id = ?"
	stmt, e := mysql_conn.MySQL.Prepare(sql)
	if e != nil {
		return e
	}
	result, e := stmt.Exec(id)
	if e != nil {
		return e
	}
	if i, e := result.RowsAffected();e==nil{
		if i>0{
			return nil
		}
	}
	return errors.New("del error")
}

func (p *ProductManger) Update(product *datamodels.Product) error {
	sql := "UPDATA product SET product_name=?,product_num=?,product_image=?,product_url=?"
	stmt, e := mysql_conn.MySQL.Prepare(sql)
	if e !=nil {
		return e
	}
	result, e := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImg, product.ProductUrl)
	if e != nil {
		return e
	}
	if i, e := result.RowsAffected();e==nil{
		if i>0 {
			return nil
		}
	}
	return errors.New("update error")
}

func (p *ProductManger) SelectByKey(id int64) (*datamodels.Product,error) {
	sql := "SELECT id,product_name,product_num,product_image,product_url FROM product WHERE id = ?"
	stmt, e := mysql_conn.MySQL.Prepare(sql)
	if e != nil {
		return nil,e
	}
	product := &datamodels.Product{}
	e = stmt.QueryRow(id).Scan(&product.Id, &product.ProductName, &product.ProductImg, &product.ProductUrl)
	if e!=nil{
		return nil,e
	}
	return product,nil
}

func (p *ProductManger) SelectAll() ([]*datamodels.Product,error) {
	sql := "SELECT id,product_name,product_num,product_image,product_url FROM product WHERE id = ?"
	stmt, e := mysql_conn.MySQL.Prepare(sql)
	if e != nil {
		return nil,e
	}
	var product []*datamodels.Product
	rows, e := stmt.Query()
	if e !=nil {
		return nil,e
	}

	for rows.Next() {
		i := &datamodels.Product{}
		e := rows.Scan(&i.Id, &i.ProductName, &i.ProductImg, &i.ProductUrl)
		if e != nil {
			return nil,e
		}
		product = append(product,i)
	}
	return product,nil
}