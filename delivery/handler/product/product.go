package product

import (
	"fmt"
	"group-project/limamart/delivery/helper"
	_middlewares "group-project/limamart/delivery/middlewares"
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
		var param _entities.Product

	err := c.Bind(&param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
		products, err := uh.productUseCase.CreateProduct(param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success create product", products))
	}
}

func (uh *ProductHandler) UpdateProductHandler() echo.HandlerFunc {
	
	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		fmt.Println("id token", idToken)
		var param _entities.Product
		id, _ := strconv.Atoi(c.Param("id"))

		getid, err := uh.productUseCase.GetProductById(id)
		
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}


		fmt.Println("id param user id", getid.UserID)

		if uint(idToken) != getid.UserID {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized2"))
		}

		
		err = c.Bind(&param)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
		products, err := uh.productUseCase.UpdateProduct(id, param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success update data product", products))
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

		getid, err := uh.productUseCase.GetProductById(id)
		
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}

		fmt.Println("id param user id", getid.UserID)

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
		
		product, err := uh.productUseCase.GetProductById(id)
		
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get product by id", product))
	}
}