package product

import (
	_entities "group-project/limamart/entities"
	_productRepository "group-project/limamart/repository/product"
)

type ProductUseCase struct {
	productRepository _productRepository.ProductRepositoryInterface
}

func NewProductUseCase(productRepo _productRepository.ProductRepositoryInterface) ProductUseCaseInterface {
	return &ProductUseCase{
		productRepository: productRepo,
	}
}

func (uuc *ProductUseCase) GetAll() ([]_entities.Product, error) {
	products, err := uuc.productRepository.GetAll()
	return products, err
}

func (uuc *ProductUseCase) CreateProduct(request _entities.Product) (_entities.Product, error) {
	books, err := uuc.productRepository.CreateProduct(request)
	return books, err
}

func (uuc *ProductUseCase) UpdateProduct(id int, request _entities.Product) (_entities.Product, error) {
	books, err := uuc.productRepository.UpdateProduct(id, request)
	return books, err
}

func (uuc *ProductUseCase) DeleteProduct(id int) error {
	err := uuc.productRepository.DeleteProduct(id)
	return err
}



func (uuc *ProductUseCase) GetProductById(id int) (_entities.Product, error) {
	books, err := uuc.productRepository.GetProductById(id)
	return books, err
}