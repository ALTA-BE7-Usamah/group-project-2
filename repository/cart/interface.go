package cart

import (
	_entities "group-project/limamart/entities"
)

type CartRepositoryInterface interface {
	GetAll(idToken int) ([]_entities.Cart, int, error)
	GetCartById(id int) (_entities.Cart, int, error)
	CreateCart(request _entities.Cart) (_entities.Cart, error)
	UpdateCart(request _entities.Cart) (_entities.Cart, int, error)
	DeleteCart(id int) error
}
