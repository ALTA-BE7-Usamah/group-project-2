package cart

import (
	"group-project/limamart/delivery/helper"
	_cartUseCase "group-project/limamart/usecase/cart"
	"net/http"
	"strconv"

	_middlewares "group-project/limamart/delivery/middlewares"
	_entities "group-project/limamart/entities"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartUseCase _cartUseCase.CartUseCaseInterface
}

func NewCartHandler(u _cartUseCase.CartUseCaseInterface) CartHandler {
	return CartHandler{
		cartUseCase: u,
	}
}

func (uh *CartHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		carts, rows, err := uh.cartUseCase.GetAll(idToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}

		responseCarts := []map[string]interface{}{}
		for i := 0; i < len(carts); i++ {
			response := map[string]interface{}{
				"id":         carts[i].ID,
				"user_id":    carts[i].UserID,
				"product_id": carts[i].ProductID,
				"qty":        carts[i].Qty,
				"sub_total":  carts[i].SubTotal,
				"product": map[string]interface{}{
					"user_id":       carts[i].Product.UserID,
					"catagory_id":   carts[i].Product.CatagoryID,
					"product_title": carts[i].Product.ProductTitle,
					"product_desc":  carts[i].Product.ProductDesc,
					"price":         carts[i].Product.Price,
					"stock":         carts[i].Product.Stock,
					"url_product":   carts[i].Product.UrlProduct},
			}
			responseCarts = append(responseCarts, response)
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all carts", responseCarts))
	}
}

func (uh *CartHandler) CreateCartHandler() echo.HandlerFunc {

	return func(c echo.Context) error {
		var param _entities.Cart

		errBind := c.Bind(&param)
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		param.UserID = uint(idToken)
		if errBind != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errBind.Error()))
		}
		_, err := uh.cartUseCase.CreateCart(param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success create cart"))
	}
}

func (uh *CartHandler) UpdateCartHandler() echo.HandlerFunc {

	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}

		var param _entities.Cart
		err = c.Bind(&param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data. please check your data"))
		}

		carts, rows, err := uh.cartUseCase.UpdateCart(id, uint(idToken), param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}

		responseCart := map[string]interface{}{
			"id":         carts.ID,
			"user_id":    carts.UserID,
			"product_id": carts.ProductID,
			"qty":        carts.Qty,
			"sub_total":  carts.SubTotal,
			"product": map[string]interface{}{
				"user_id":       carts.Product.UserID,
				"catagory_id":   carts.Product.CatagoryID,
				"product_title": carts.Product.ProductTitle,
				"product_desc":  carts.Product.ProductDesc,
				"price":         carts.Product.Price,
				"stock":         carts.Product.Stock,
				"url_product":   carts.Product.UrlProduct},
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success update cart", responseCart))
	}
}

func (uh *CartHandler) DeleteCartHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		id, _ := strconv.Atoi(c.Param("id"))

		getid, rows, err := uh.cartUseCase.GetCartById(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}

		if uint(idToken) != getid.UserID {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		err = uh.cartUseCase.DeleteCart(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success delete cart"))
	}
}
