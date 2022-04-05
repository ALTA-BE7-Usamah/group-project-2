package user

import (
	_entities "group-project/limamart/entities"
)

type UserRepositoryInterface interface {
	CreateUser(request _entities.User) (_entities.User, error)
}