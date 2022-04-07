package order

import (
	"errors"
	_entities "group-project/limamart/entities"
	_cartRepository "group-project/limamart/repository/cart"
	_orderRepository "group-project/limamart/repository/order"
)

type OrderUseCase struct {
	orderRepository _orderRepository.OrderRepositoryInterface
	cartRepository  _cartRepository.CartRepositoryInterface
}

func NewOrderUseCase(orderRepo _orderRepository.OrderRepositoryInterface, cartRepo _cartRepository.CartRepositoryInterface) OrderUseCaseInterface {
	return &OrderUseCase{
		orderRepository: orderRepo,
		cartRepository:  cartRepo,
	}
}

func (uuc *OrderUseCase) GetAll(id int) ([]_entities.Order, error) {
	carts, err := uuc.orderRepository.GetAll(id)
	return carts, err
}

func (uuc *OrderUseCase) CreateOrder(creatOrder _entities.Order, orderCartID []uint, idToken uint) (_entities.Order, int, error) {
	carts, rows, err := uuc.cartRepository.GetAll(int(idToken))
	if rows == 0 {
		return creatOrder, 0, errors.New("failed get all cart")
	}
	if err != nil {
		return creatOrder, 0, err
	}
	for i := 0; i < len(carts); i++ {
		creatOrder.TotalPrice += carts[i].SubTotal
	}

	creatOrder.StatusOrder = "purchased"

	order, rows, err := uuc.orderRepository.CreateOrder(creatOrder, orderCartID)
	return order, rows, err
}
