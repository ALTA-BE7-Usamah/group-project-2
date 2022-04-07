package order

import (
	_entities "group-project/limamart/entities"
)

type OrderRepositoryInterface interface {
	GetAllOrder(idToken int) ([]_entities.OrdersDetail, int, error)
	CreateOrder(creatOrder _entities.Order, orderCartID []uint) (_entities.Order, int, error)
	CancelOrder(cancelOrder _entities.OrdersDetail) (_entities.OrdersDetail, int, error)
}
