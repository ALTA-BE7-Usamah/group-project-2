package order

import (
	"group-project/limamart/delivery/helper"
	_orderUseCase "group-project/limamart/usecase/order"
	"net/http"
	"strconv"

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

func (uh *OrderHandler) GetAllOrdersHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		orders, products, rows, err := uh.orderUseCase.GetAllOrder(idToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}

		responseOrders := []map[string]interface{}{}
		for i := 0; i < len(orders); i++ {
			response := map[string]interface{}{
				"id":          orders[i].ID,
				"user_id":     orders[i].UserID,
				"product_id":  orders[i].ProductID,
				"total_price": orders[i].TotalPrice,
				"status":      orders[i].Status,
				"product": map[string]interface{}{
					"id":            products[i].ID,
					"user_id":       products[i].UserID,
					"catagory_id":   products[i].CatagoryID,
					"product_title": products[i].ProductTitle,
					"price":         products[i].Price,
					"stock":         products[i].Stock,
					"url_product":   products[i].UrlProduct},
			}

			responseOrders = append(responseOrders, response)
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all orders", responseOrders))
	}
}

func (uh *OrderHandler) CreateOrderHandler() echo.HandlerFunc {

	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		var orderRequest helper.OrderRequestFormat
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

func (uh *OrderHandler) CancelOrderHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}

		var cancelOrder _entities.OrdersDetail
		errBind := c.Bind(&cancelOrder)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data. please check your data"))
		}

		_, rows, err := uh.orderUseCase.CancelOrder(cancelOrder, uint(id), uint(idToken))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("successfully cancel order"))
	}
}
