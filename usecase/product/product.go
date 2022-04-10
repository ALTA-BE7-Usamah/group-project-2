package product

import (
	"errors"
	_entities "group-project/limamart/entities"
	_cartRepository "group-project/limamart/repository/cart"
	_productRepository "group-project/limamart/repository/product"
)

type ProductUseCase struct {
	productRepository _productRepository.ProductRepositoryInterface
	cartRepository    _cartRepository.CartRepositoryInterface
}

func NewProductUseCase(productRepo _productRepository.ProductRepositoryInterface, cartRepo _cartRepository.CartRepositoryInterface) ProductUseCaseInterface {
	return &ProductUseCase{
		productRepository: productRepo,
		cartRepository:    cartRepo,
	}
}

func (uuc *ProductUseCase) GetAll() ([]_entities.Product, error) {
	products, err := uuc.productRepository.GetAll()
	return products, err
}

func (uuc *ProductUseCase) CreateProduct(request _entities.Product) (_entities.Product, error) {
	product, err := uuc.productRepository.CreateProduct(request)
	if request.CatagoryID == 0 {
		return product, errors.New("can't be empty")
	}
	if request.ProductTitle == "" {
		return product, errors.New("can't be empty")
	}
	if request.ProductDesc == "" {
		return product, errors.New("can't be empty")
	}
	if request.Price == 0 {
		return product, errors.New("can't be empty")
	}
	if request.Stock == 0 {
		return product, errors.New("can't be empty")
	}
	if request.UrlProduct == "" {
		return product, errors.New("can't be empty")
	}
	return product, err
}

func (uuc *ProductUseCase) UpdateProduct(request _entities.Product, id uint, idToken uint) (_entities.Product, int, error) {
	productFind, rows, err := uuc.productRepository.GetProductById(int(id))
	if err != nil {
		return productFind, 0, err
	}
	if rows == 0 {
		return productFind, 0, nil
	}
	if productFind.UserID != idToken {
		return productFind, 1, errors.New("unauthorized")
	}
	if request.CatagoryID != 0 {
		productFind.CatagoryID = request.CatagoryID
	}
	if request.ProductTitle != "" {
		productFind.ProductTitle = request.ProductTitle
	}
	if request.ProductDesc != "" {
		productFind.ProductDesc = request.ProductDesc
	}
	if request.Price != 0 {
		productFind.Price = request.Price
	}
	if request.Stock != 0 {
		productFind.Stock = request.Stock
	}
	if request.UrlProduct != "" {
		productFind.UrlProduct = request.UrlProduct
	}
	product, rows, err := uuc.productRepository.UpdateProduct(productFind)
	return product, rows, err
}

func (uuc *ProductUseCase) DeleteProduct(id int) error {
	product, _, err := uuc.productRepository.GetProductById(id)
	if err != nil {
		return err
	}

	cart, _, err := uuc.cartRepository.GetCartByProductId(int(product.ID))
	if err != nil {
		return err
	}
	errDelete := uuc.productRepository.DeleteProduct(id, cart)
	return errDelete
}

func (uuc *ProductUseCase) GetProductById(id int) (_entities.Product, int, error) {
	product, rows, err := uuc.productRepository.GetProductById(id)
	return product, rows, err
}

func (uuc *ProductUseCase) GetAllProductUser(userID uint) ([]_entities.Product, error) {
	products, err := uuc.productRepository.GetAllProductUser(userID)
	return products, err
}
