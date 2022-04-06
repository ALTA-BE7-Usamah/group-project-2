package product

import (
	_entities "group-project/limamart/entities"
)

type ProductRepositoryInterface interface {
	GetAll() ([]_entities.Product, error)
	CreateProduct(request _entities.Product) (_entities.Product, error)
	UpdateProduct(request _entities.Product) (_entities.Product, int, error)
	DeleteProduct(id int) error
	GetProductById(id int) (_entities.Product, int, error)
}
