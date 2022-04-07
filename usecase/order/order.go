package order

import (
	"errors"
	_entities "group-project/limamart/entities"
	_cartRepository "group-project/limamart/repository/cart"
	_orderRepository "group-project/limamart/repository/order"
	_productRepository "group-project/limamart/repository/product"
)

type OrderUseCase struct {
	orderRepository   _orderRepository.OrderRepositoryInterface
	cartRepository    _cartRepository.CartRepositoryInterface
	productRepository _productRepository.ProductRepositoryInterface
}

func NewOrderUseCase(orderRepo _orderRepository.OrderRepositoryInterface, cartRepo _cartRepository.CartRepositoryInterface, productRepo _productRepository.ProductRepositoryInterface) OrderUseCaseInterface {
	return &OrderUseCase{
		orderRepository:   orderRepo,
		cartRepository:    cartRepo,
		productRepository: productRepo,
	}
}

func (uuc *OrderUseCase) GetAllOrder(idToken int) ([]_entities.OrdersDetail, int, error) {
	carts, rows, err := uuc.orderRepository.GetAllOrder(idToken)
	return carts, rows, err
}

func (uuc *OrderUseCase) CreateOrder(creatOrder _entities.Order, orderCartID []uint, idToken uint) (_entities.Order, int, error) {
	for i := 0; i < len(orderCartID); i++ {
		carts, rows, err := uuc.cartRepository.GetCartById(int(orderCartID[i]))
		if rows == 0 {
			return creatOrder, 0, errors.New("failed get all cart")
		}
		if err != nil {
			return creatOrder, 0, err
		}
		creatOrder.TotalPrice += carts.SubTotal
	}

	creatOrder.StatusOrder = "purchased"

	order, rows, err := uuc.orderRepository.CreateOrder(creatOrder, orderCartID)
	return order, rows, err
}
