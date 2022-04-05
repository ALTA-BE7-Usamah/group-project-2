package user

import (
	_entities "group-project/limamart/entities"
	_userRepository "group-project/limamart/repository/users"
)

type UserUseCase struct {
	userRepository _userRepository.UserRepositoryInterface
}

func NewUserUseCase(userRepo _userRepository.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		userRepository: userRepo,
	}
}

func (uuc *UserUseCase) CreateUser(request _entities.User) (_entities.User, error) {
	users, err := uuc.userRepository.CreateUser(request)
	return users, err
}