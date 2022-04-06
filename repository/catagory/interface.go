package catagory

import (
	_entities "group-project/limamart/entities"
)

type CatagoryRepositoryInterface interface {
	GetAllCatagory() ([]_entities.Catagory, error)
}
