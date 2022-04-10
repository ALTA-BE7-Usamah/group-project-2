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

func (ur *ProductRepository) GetProductById(id int) (_entities.Product, int, error) {
	var products _entities.Product
	tx := ur.DB.Unscoped().Find(&products, id)
	if tx.Error != nil {
		return products, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return products, 0, nil
	}
	return products, int(tx.RowsAffected), nil
}

func (ur *ProductRepository) CreateProduct(request _entities.Product) (_entities.Product, error) {
	yx := ur.DB.Save(&request)
	if yx.Error != nil {
		return request, yx.Error
	}
	return request, nil
}

func (ur *ProductRepository) UpdateProduct(request _entities.Product) (_entities.Product, int, error) {
	tx := ur.DB.Save(&request)
	if tx.Error != nil {
		return request, 0, tx.Error
	}
	return request, int(tx.RowsAffected), nil
}

func (ur *ProductRepository) DeleteProduct(id int, cart []_entities.Cart) error {
	for i := 0; i < len(cart); i++ {
		var cartDelete _entities.Cart
		errCart := ur.DB.Unscoped().Delete(cartDelete, cart[i].ProductID).Error
		if errCart != nil {
			return errCart
		}
	}

	err := ur.DB.Delete(&_entities.Product{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *ProductRepository) GetAllProductUser(userID uint) ([]_entities.Product, error) {
	var products []_entities.Product
	tx := ur.DB.Where("user_id = ?", userID).Find(&products)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return products, nil
}
