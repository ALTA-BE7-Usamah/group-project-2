package user

import (
	_entities "group-project/limamart/entities"
)

type UserRepositoryInterface interface {
	GetUserById(id int) (_entities.User, int, error)
	CreateUser(request _entities.User) (_entities.User, error)
	UpdateUser(request _entities.User) (_entities.User, int, error)
	DeleteUser(id int) error
}
