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
		var param _entities.Order

	err := c.Bind(&param)
	idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		param.UserID = uint(idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
		orders, err := uh.orderUseCase.CreateOrder(param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success create order", orders))
	}
}