package cart

import (
	_entities "group-project/limamart/entities"
)

type CartUseCaseInterface interface {
	GetAll(id int) ([]_entities.Cart, error)
	GetCartById(id int) (_entities.Cart, error)
	CreateCart(request _entities.Cart) (_entities.Cart, error)
	UpdateCart(id int, request _entities.Cart) (_entities.Cart, error)
	DeleteCart(id int) error
}