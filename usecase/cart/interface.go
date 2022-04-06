package cart

import (
	_entities "group-project/limamart/entities"
)

type CartUseCaseInterface interface {
	GetAll(idToken int) ([]_entities.Cart, int, error)
	GetCartById(id int) (_entities.Cart, int, error)
	CreateCart(request _entities.Cart) (_entities.Cart, error)
	UpdateCart(id int, idToken uint, request _entities.Cart) (_entities.Cart, int, error)
	DeleteCart(id int) error
}
