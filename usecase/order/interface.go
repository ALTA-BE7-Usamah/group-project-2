package order

import (
	_entities "group-project/limamart/entities"
)

type OrderUseCaseInterface interface {
	GetAllOrder(idToken int) ([]_entities.OrdersDetail, int, error)
	CreateOrder(request _entities.Order, orderCartID []uint, idToken uint) (_entities.Order, int, error)
}
