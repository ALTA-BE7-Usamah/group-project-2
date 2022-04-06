package cart

import (
	_entities "group-project/limamart/entities"

	"gorm.io/gorm"
)

type CartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		DB: db,
	}

}

func (ur *CartRepository) GetAll(idToken int) ([]_entities.Cart, int, error) {
	var carts []_entities.Cart
	tx := ur.DB.Preload("Product").Where("user_id = ?", idToken).Find(&carts)
	if tx.Error != nil {
		return carts, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return carts, 0, nil
	}
	return carts, int(tx.RowsAffected), nil
}

func (ur *CartRepository) GetCartById(id int) (_entities.Cart, error) {
	var carts _entities.Cart
	tx := ur.DB.Find(&carts, id)
	if tx.Error != nil {
		return carts, tx.Error
	}

	return carts, nil
}

func (ur *CartRepository) CreateCart(request _entities.Cart) (_entities.Cart, error) {

	yx := ur.DB.Save(&request)
	if yx.Error != nil {
		return request, yx.Error
	}

	return request, nil
}

func (ur *CartRepository) UpdateCart(id int, request _entities.Cart) (_entities.Cart, error) {
	err := ur.DB.Where("id = ?", id).Updates(&request).Error
	// err := ur.DB.Model(&_entities.Cart{}).Where("id = ?", id).Updates(&request).Error
	if err != nil {
		return request, err
	}

	return request, nil
}

func (ur *CartRepository) DeleteCart(id int) error {

	err := ur.DB.Unscoped().Delete(&_entities.Cart{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
