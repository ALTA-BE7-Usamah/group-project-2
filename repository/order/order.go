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

func (ur *OrderRepository) GetAllOrder(idToken int) ([]_entities.OrdersDetail, int, error) {
	var orders []_entities.OrdersDetail
	tx := ur.DB.Preload("Product").Where("user_id = ?", idToken).Find(&orders)
	if tx.Error != nil {
		return orders, 0, tx.Error
	}
	return orders, int(tx.RowsAffected), nil
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
		ordersDetail.ProductID = carts.ProductID
		ordersDetail.TotalPrice = carts.SubTotal
		yx := ur.DB.Save(&ordersDetail)
		if yx.Error != nil {
			return creatOrder, 0, yx.Error
		}
		var products _entities.Product
		zx := ur.DB.Find(&products, carts.ProductID)
		if zx.Error != nil {
			return creatOrder, 0, zx.Error
		}
		products.Stock -= carts.Qty
		ux := ur.DB.Save(&products)
		if ux.Error != nil {
			return creatOrder, 0, ux.Error
		}

		err := ur.DB.Unscoped().Delete(&_entities.Cart{}, orderCartID[i]).Error
		if err != nil {
			return creatOrder, 0, err
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

func (ur *OrderRepository) CancelOrder(cancelOrder _entities.OrdersDetail) (_entities.OrdersDetail, int, error) {
	tx := ur.DB.Save(&cancelOrder)
	if tx.Error != nil {
		return cancelOrder, 0, tx.Error
	}
	return cancelOrder, int(tx.RowsAffected), nil
}
