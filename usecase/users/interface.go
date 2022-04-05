package user

import (
	_entities "group-project/limamart/entities"
)

type UserUseCaseInterface interface {
	CreateUser(request _entities.User) (_entities.User, error)
}