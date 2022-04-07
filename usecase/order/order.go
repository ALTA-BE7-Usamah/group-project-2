package order

import (
	"errors"
	_entities "group-project/limamart/entities"
	_cartRepository "group-project/limamart/repository/cart"
	_orderRepository "group-project/limamart/repository/order"
	_productRepository "group-project/limamart/repository/product"
)

type OrderUseCase struct {
	orderRepository   _orderRepository.OrderRepositoryInterface
	cartRepository    _cartRepository.CartRepositoryInterface
	productRepository _productRepository.ProductRepositoryInterface
}

func NewOrderUseCase(orderRepo _orderRepository.OrderRepositoryInterface, cartRepo _cartRepository.CartRepositoryInterface, productRepo _productRepository.ProductRepositoryInterface) OrderUseCaseInterface {
	return &OrderUseCase{
		orderRepository:   orderRepo,
		cartRepository:    cartRepo,
		productRepository: productRepo,
	}
}

func (uuc *OrderUseCase) GetAllOrder(idToken int) ([]_entities.OrdersDetail, int, error) {
	order, rows, err := uuc.orderRepository.GetAllOrder(idToken)
	return order, rows, err
}

func (uuc *OrderUseCase) CreateOrder(creatOrder _entities.Order, orderCartID []uint, idToken uint) (_entities.Order, int, error) {
	for i := 0; i < len(orderCartID); i++ {
		carts, rows, err := uuc.cartRepository.GetCartById(int(orderCartID[i]))
		if rows == 0 {
			return creatOrder, 0, errors.New("failed get all cart")
		}
		if err != nil {
			return creatOrder, 0, err
		}
		creatOrder.TotalPrice += carts.SubTotal
	}

	creatOrder.StatusOrder = "purchased"

	order, rows, err := uuc.orderRepository.CreateOrder(creatOrder, orderCartID)
	return order, rows, err
}

func (uuc *OrderUseCase) GetHistoriOrderbyID(id int) (_entities.OrdersDetail, int, error) {
	order, rows, err := uuc.orderRepository.GetHistoriOrderbyID(id)
	return order, rows, err
}

func (uuc *OrderUseCase) CancelOrder(cancelOrder _entities.OrdersDetail, id uint, idToken uint) (_entities.OrdersDetail, int, error) {
	order, rows, err := uuc.orderRepository.GetHistoriOrderbyID(int(id))
	if err != nil {
		return order, 0, err
	}
	if rows == 0 {
		return order, 0, nil
	}
	if order.UserID != idToken {
		return order, 1, errors.New("unauthorized")
	}
	if cancelOrder.UserID != 0 {
		order.UserID = cancelOrder.UserID
	}
	if cancelOrder.ProductID != 0 {
		order.ProductID = cancelOrder.ProductID
	}
	if cancelOrder.TotalPrice != 0 {
		order.TotalPrice = cancelOrder.TotalPrice
	}
	order.Status = "canceled"
	orderUpdate, rowsUpdate, errUpdate := uuc.orderRepository.CancelOrder(order)
	return orderUpdate, rowsUpdate, errUpdate
}
