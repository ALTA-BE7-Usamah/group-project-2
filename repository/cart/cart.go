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

func (ur *CartRepository) GetCartById(id int) (_entities.Cart, int, error) {
	var carts _entities.Cart
	tx := ur.DB.Preload("Product").Find(&carts, id)
	if tx.Error != nil {
		return carts, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return carts, 0, nil
	}
	return carts, int(tx.RowsAffected), nil
}

func (ur *CartRepository) GetCartByProductId(idProduct int) ([]_entities.Cart, int, error) {
	var cart []_entities.Cart
	tx := ur.DB.Where("product_id = ?", idProduct).Find(&cart)
	if tx.Error != nil {
		return cart, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return cart, 0, nil
	}
	return cart, int(tx.RowsAffected), nil
}

func (ur *CartRepository) CreateCart(request _entities.Cart) (_entities.Cart, error) {
	yx := ur.DB.Save(&request)
	if yx.Error != nil {
		return request, yx.Error
	}
	return request, nil
}

func (ur *CartRepository) UpdateCart(request _entities.Cart) (_entities.Cart, int, error) {
	tx := ur.DB.Save(&request)
	if tx.Error != nil {
		return request, 0, tx.Error
	}
	return request, int(tx.RowsAffected), nil
}

func (ur *CartRepository) DeleteCart(id int) error {

	err := ur.DB.Unscoped().Delete(&_entities.Cart{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
