package product

import (
	_entities "group-project/limamart/entities"
	_catagoryRepository "group-project/limamart/repository/catagory"
)

type CatagoryUseCase struct {
	catagoryRepository _catagoryRepository.CatagoryRepositoryInterface
}

func NewCatagoryUseCase(catagoryRepo _catagoryRepository.CatagoryRepositoryInterface) CatagoryUseCaseInterface {
	return &CatagoryUseCase{
		catagoryRepository: catagoryRepo,
	}
}

func (cuc *CatagoryUseCase) GetAllCatagory() ([]_entities.Catagory, error) {
	catagory, err := cuc.catagoryRepository.GetAllCatagory()
	return catagory, err
}
