package order

import (
	_entities "group-project/limamart/entities"
)

type OrderRepositoryInterface interface {
	GetAll(id int) ([]_entities.Order, error)
	CreateOrder(creatOrder _entities.Order, orderCartID []uint) (_entities.Order, int, error)
}
