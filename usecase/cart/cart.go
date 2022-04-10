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

func (uuc *CartUseCase) GetCartById(id int) (_entities.Cart, int, error) {
	carts, rows, err := uuc.cartRepository.GetCartById(id)
	return carts, rows, err
}

func (uuc *CartUseCase) GetCartByProductId(idProduct int) ([]_entities.Cart, int, error) {
	carts, rows, err := uuc.cartRepository.GetCartByProductId(idProduct)
	return carts, rows, err
}

func (uuc *CartUseCase) CreateCart(request _entities.Cart, idToken uint) (_entities.Cart, error) {
	products, rows, err := uuc.productRepository.GetProductById(int(request.ProductID))
	if rows == 0 {
		return request, errors.New("failed product by id")
	}
	if err != nil {
		return request, err
	}
	if request.ProductID == 0 {
		return request, errors.New("can't be empty")
	}
	if request.Qty == 0 {
		return request, errors.New("can't be empty")
	}
	if request.Qty > products.Stock {
		return _entities.Cart{}, errors.New("this product is out of stock")
	}
	request.SubTotal = request.Qty * products.Price

	carts, rows, err := uuc.cartRepository.GetAll(int(idToken))
	for i := 0; i < len(carts); i++ {
		if carts[i].ProductID == request.ProductID {
			if carts[i].Qty+request.Qty > products.Stock {
				return _entities.Cart{}, errors.New("this product is out of stock")
			}
			carts[i].Qty += request.Qty
			carts[i].SubTotal += (request.Qty * products.Price)
			cart, err := uuc.cartRepository.CreateCart(carts[i])
			return cart, err
		}
	}

	cartCreate, errCreate := uuc.cartRepository.CreateCart(request)
	return cartCreate, errCreate
}

func (uuc *CartUseCase) UpdateCart(id int, idToken uint, request _entities.Cart) (_entities.Cart, int, error) {
	cart, rows, err := uuc.cartRepository.GetCartById(id)
	if err != nil {
		return request, 0, err
	}
	if rows == 0 {
		return request, 0, nil
	}
	if cart.UserID != idToken {
		return request, 1, errors.New("unauthorized")
	}
	products, rows, err := uuc.productRepository.GetProductById(int(cart.ProductID))
	if err != nil {
		return request, 0, err
	}
	if rows == 0 {
		return request, 0, nil
	}
	if request.ProductID != 0 {
		cart.ProductID = request.ProductID
	}
	if request.Qty != 0 {
		cart.Qty = request.Qty
	}
	if request.Qty > products.Stock {
		return request, 0, errors.New("this product is out of stock")
	}

	cart.SubTotal = cart.Qty * products.Price

	carts, rows, err := uuc.cartRepository.UpdateCart(cart)
	return carts, rows, err
}

func (uuc *CartUseCase) DeleteCart(id int) error {
	err := uuc.cartRepository.DeleteCart(id)
	return err
}
