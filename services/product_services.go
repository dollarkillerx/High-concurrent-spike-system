/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-17
* Time: 下午5:46
* */
package services

import (
	"High-concurrent-spike-system/datamodels"
	"High-concurrent-spike-system/repositories"
)

type IProductService interface {
	GetProductById(int64) (*datamodels.Product,error)
	GetAllProduct()([]*datamodels.Product,error)
	DeleteProductById(int64) error
	InsertProduct(*datamodels.Product)(int64,error)
	UpdateProduct(*datamodels.Product) error
}

type ProductService struct {
	productRepository repositories.IProduct
}

func NewProductService(product repositories.IProduct) IProductService {
	return &ProductService{
		productRepository:product,
	}
}

func (p *ProductService) GetProductById(id int64) (*datamodels.Product,error) {
	product, e := p.productRepository.SelectByKey(id)
	return product,e
}

func (p *ProductService) GetAllProduct()([]*datamodels.Product,error) {
	products, e := p.productRepository.SelectAll()
	return products,e
}

func (p *ProductService) DeleteProductById(id int64) error {
	e := p.productRepository.Delete(id)
	return e
}

func (p *ProductService) InsertProduct(data *datamodels.Product)(int64,error) {
	i, e := p.productRepository.Insert(data)
	return i,e
}

func (p *ProductService) UpdateProduct(data *datamodels.Product) error {
	e := p.productRepository.Update(data)
	return e
}


