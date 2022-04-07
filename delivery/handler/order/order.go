package order

import (
	"fmt"
	"group-project/limamart/delivery/helper"
	_orderUseCase "group-project/limamart/usecase/order"
	"net/http"

	_middlewares "group-project/limamart/delivery/middlewares"
	_entities "group-project/limamart/entities"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderUseCase _orderUseCase.OrderUseCaseInterface
}

func NewOrderHandler(u _orderUseCase.OrderUseCaseInterface) OrderHandler {
	return OrderHandler{
		orderUseCase: u,
	}
}

func (uh *OrderHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		fmt.Println("id token", idToken)

		orders, err := uh.orderUseCase.GetAll(idToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all orders", orders))
	}
}

func (uh *OrderHandler) CreateOrderHandler() echo.HandlerFunc {

	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		var orderRequest _entities.OrderRequestFormat
		err := c.Bind(&orderRequest)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}

		address := _entities.Address{
			UserID:  uint(idToken),
			Street:  orderRequest.Address.Street,
			City:    orderRequest.Address.City,
			State:   orderRequest.Address.State,
			ZipCode: orderRequest.Address.ZipCode,
		}

		creditcard := _entities.CreditCard{
			Type:   orderRequest.CreditCard.Type,
			Name:   orderRequest.CreditCard.Name,
			Number: orderRequest.CreditCard.Number,
			CVV:    orderRequest.CreditCard.CVV,
			Month:  orderRequest.CreditCard.Month,
			Year:   orderRequest.CreditCard.Year,
		}

		creatOrder := _entities.Order{
			UserID:     uint(idToken),
			Address:    address,
			CreditCard: creditcard,
		}

		var orderCartID = orderRequest.CartID

		_, rows, err := uh.orderUseCase.CreateOrder(creatOrder, orderCartID, uint(idToken))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success create order"))
	}
}
