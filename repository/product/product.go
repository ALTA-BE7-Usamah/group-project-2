package product

import (
	_entities "group-project/limamart/entities"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (ur *ProductRepository) GetAll() ([]_entities.Product, error) {
	var products []_entities.Product
	tx := ur.DB.Find(&products)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return products, nil
}

func (ur *ProductRepository) GetProductById(id int) (_entities.Product, error) {
	var products _entities.Product
	tx := ur.DB.Find(&products, id)
	if tx.Error != nil {
		return products, tx.Error
	}

	return products, nil
}

func (ur *ProductRepository) CreateProduct(request _entities.Product) (_entities.Product, error) {
	yx := ur.DB.Save(&request)
	if yx.Error != nil {
		return request , yx.Error
	}

	return request, nil
}

func (ur *ProductRepository) UpdateProduct(id int, request _entities.Product) (_entities.Product, error) {
	err := ur.DB.Model(&_entities.Product{}).Where("id = ?", id).Updates(request).Error
	if err != nil {
		return request , err
	}

	return request, nil
}

func (ur *ProductRepository) DeleteProduct(id int) error {
	
	err := ur.DB.Unscoped().Delete(&_entities.Product{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

