package order

import (
	_entities "group-project/limamart/entities"

	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		DB: db,
	}

}

func (ur *OrderRepository) GetAll(id int) ([]_entities.Order, error) {
	var orders []_entities.Order
	tx := ur.DB.Where("user_id = ?", id).Find(&orders)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return orders, nil
}

func (ur *OrderRepository) CreateOrder(request _entities.Order, cart []_entities.Cart) (_entities.Order, int, error) {
	yx := ur.DB.Save(&request)
	if yx.Error != nil {
		return request, 0, yx.Error
	}
	tx := ur.DB.Save(&cart)
	if tx.Error != nil {
		return request, 0, tx.Error
	}
	return request, int(tx.RowsAffected), nil
}
