package product

import (
	_entities "group-project/limamart/entities"
)

type CatagoryUseCaseInterface interface {
	GetAllCatagory() ([]_entities.Catagory, error)
}
