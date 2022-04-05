package user

import (
	_entities "group-project/limamart/entities"
)

type UserUseCaseInterface interface {
	CreateUser(request _entities.User) (_entities.User, error)
	UpdateUser(id int, request _entities.User) (_entities.User, error)
	DeleteUser(id int) error
	GetUserById(id int) (_entities.User, error)
}
