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

func (ur *OrderRepository) CreateOrder(creatOrder _entities.Order, orderCartID []uint) (_entities.Order, int, error) {
	yx := ur.DB.Save(&creatOrder)
	if yx.Error != nil {
		return creatOrder, 0, yx.Error
	}
	for i := 0; i < len(orderCartID); i++ {
		var ordersDetail _entities.OrdersDetail
		var carts _entities.Cart
		tx := ur.DB.Preload("Product").Find(&carts, orderCartID[i])
		if tx.Error != nil {
			return creatOrder, 0, tx.Error
		}
		ordersDetail.UserID = creatOrder.UserID
		ordersDetail.OrderID = creatOrder.ID
		ordersDetail.ProductID = carts.ProductID
		ordersDetail.TotalPrice = carts.SubTotal
		yx := ur.DB.Save(&ordersDetail)
		if yx.Error != nil {
			return creatOrder, 0, yx.Error
		}
	}
	zx := ur.DB.Save(&creatOrder.Address)
	if zx.Error != nil {
		return creatOrder, 0, zx.Error
	}
	xx := ur.DB.Save(&creatOrder.CreditCard)
	if xx.Error != nil {
		return creatOrder, 0, xx.Error
	}
	return creatOrder, int(yx.RowsAffected), nil
}
