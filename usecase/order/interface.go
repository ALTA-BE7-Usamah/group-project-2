package order

import (
	_entities "group-project/limamart/entities"
)

type OrderUseCaseInterface interface {
	GetAll(id int) ([]_entities.Order, error)
	CreateOrder(request _entities.Order, idToken uint) (_entities.Order, int, error)
}
