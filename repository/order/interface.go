package order

import (
	_entities "group-project/limamart/entities"
)

type OrderRepositoryInterface interface {
	GetAll(id int) ([]_entities.Order, error)
	CreateOrder(request _entities.Order, cart []_entities.Cart) (_entities.Order, int, error)
}
