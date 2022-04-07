package order

import (
	_entities "group-project/limamart/entities"
)

type OrderUseCaseInterface interface {
	GetAllOrder(idToken int) ([]_entities.OrdersDetail, int, error)
	CreateOrder(request _entities.Order, orderCartID []uint, idToken uint) (_entities.Order, int, error)
	GetHistoriOrderbyID(id int) (_entities.OrdersDetail, int, error)
	CancelOrder(cancelOrder _entities.OrdersDetail, id uint, idToken uint) (_entities.OrdersDetail, int, error)
}
