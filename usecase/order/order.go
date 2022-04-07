package order

import (
	_entities "group-project/limamart/entities"
	_orderRepository "group-project/limamart/repository/order"
)

type OrderUseCase struct {
	orderRepository    _orderRepository.OrderRepositoryInterface
}

func NewOrderUseCase(orderRepo _orderRepository.OrderRepositoryInterface) OrderUseCaseInterface {
	return &OrderUseCase{
		orderRepository: orderRepo,
	}
}

func (uuc *OrderUseCase) GetAll(id int) ([]_entities.Order, error) {
	carts, err := uuc.orderRepository.GetAll(id)
	return carts, err
}

func (uuc *OrderUseCase) CreateOrder(request _entities.Order) (_entities.Order, error) {
	carts, err := uuc.orderRepository.CreateOrder(request)
	return carts, err
}
