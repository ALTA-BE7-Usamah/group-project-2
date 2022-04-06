package product

import (
	"fmt"
	"group-project/limamart/delivery/helper"
	_middlewares "group-project/limamart/delivery/middlewares"
	"group-project/limamart/entities"
	_productUseCase "group-project/limamart/usecase/product"
	"net/http"
	"strconv"

	_entities "group-project/limamart/entities"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productUseCase _productUseCase.ProductUseCaseInterface
}

func NewProductHandler(u _productUseCase.ProductUseCaseInterface) ProductHandler {
	return ProductHandler{
		productUseCase: u,
	}
}

func (uh *ProductHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		products, err := uh.productUseCase.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all products", products))
	}
}

func (uh *ProductHandler) CreateProductHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		var param _entities.Product

		err := c.Bind(&param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		param.UserID = uint(idToken)

		products, err := uh.productUseCase.CreateProduct(param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success create product", products))
	}
}

func (uh *ProductHandler) UpdateProductHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		//mendapatkan id dari token yang dimasukkan
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}

		var updateProduct entities.Product
		errBind := c.Bind(&updateProduct)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data. please check your data"))
		}

		product, rows, err := uh.productUseCase.UpdateProduct(updateProduct, uint(id), uint(idToken))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success update data product", product))
	}
}

func (uh *ProductHandler) DeleteProductHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		fmt.Println("id token", idToken)

		id, _ := strconv.Atoi(c.Param("id"))

		getid, rows, err := uh.productUseCase.GetProductById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		if uint(idToken) != getid.UserID {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		err = uh.productUseCase.DeleteProduct(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success delete product", err))
	}
}

func (uh *ProductHandler) GetProductByIdHandler() echo.HandlerFunc {

	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}

		product, rows, err := uh.productUseCase.GetProductById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get product by id", product))
	}
}
