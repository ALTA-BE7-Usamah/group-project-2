package cart

import (
	"errors"
	_entities "group-project/limamart/entities"
	_cartRepository "group-project/limamart/repository/cart"
	_productRepository "group-project/limamart/repository/product"
)

type CartUseCase struct {
	cartRepository    _cartRepository.CartRepositoryInterface
	productRepository _productRepository.ProductRepositoryInterface
}

func NewCartUseCase(cartRepo _cartRepository.CartRepositoryInterface, productRepo _productRepository.ProductRepositoryInterface) CartUseCaseInterface {
	return &CartUseCase{
		cartRepository:    cartRepo,
		productRepository: productRepo,
	}
}

func (uuc *CartUseCase) GetAll(idToken int) ([]_entities.Cart, int, error) {
	carts, rows, err := uuc.cartRepository.GetAll(idToken)
	return carts, rows, err
}

func (uuc *CartUseCase) GetCartById(id int) (_entities.Cart, error) {
	carts, err := uuc.cartRepository.GetCartById(id)
	return carts, err
}

func (uuc *CartUseCase) CreateCart(request _entities.Cart) (_entities.Cart, error) {
	products, rows, err := uuc.productRepository.GetProductById(int(request.ProductID))
	if rows == 0 {
		return request, err
	}

	if request.Qty > products.Stock {
		return _entities.Cart{}, errors.New("this product is out of stock")
	}

	request.SubTotal = request.Qty * products.Price

	carts, err := uuc.cartRepository.CreateCart(request)
	return carts, err
}

func (uuc *CartUseCase) UpdateCart(id int, request _entities.Cart) (_entities.Cart, error) {
	cart, err := uuc.cartRepository.GetCartById(id)
	products, rows, err := uuc.productRepository.GetProductById(int(cart.ProductID))
	if rows == 0 {
		return cart, err
	}

	if request.Qty > products.Stock {
		return _entities.Cart{}, errors.New("this product is out of stock")
	}

	request.SubTotal = request.Qty * products.Price

	carts, err := uuc.cartRepository.UpdateCart(id, request)
	return carts, err
}

func (uuc *CartUseCase) DeleteCart(id int) error {
	err := uuc.cartRepository.DeleteCart(id)
	return err
}
