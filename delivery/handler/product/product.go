package product

import (
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

		responseProduct := []map[string]interface{}{}
		for i := 0; i < len(products); i++ {
			response := map[string]interface{}{
				"id":            products[i].ID,
				"user_id":       products[i].UserID,
				"catagory_id":   products[i].CatagoryID,
				"product_title": products[i].ProductTitle,
				"product_desc":  products[i].ProductDesc,
				"price":         products[i].Price,
				"stock":         products[i].Stock,
				"url_product":   products[i].UrlProduct,
			}
			responseProduct = append(responseProduct, response)
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all products", responseProduct))
	}
}

func (uh *ProductHandler) CreateProductHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		var param _entities.Product

		errBind := c.Bind(&param)
		if errBind != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errBind.Error()))
		}
		param.UserID = uint(idToken)

		_, err := uh.productUseCase.CreateProduct(param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success create product"))
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

		var updateProduct _entities.Product
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

		responseProduct := map[string]interface{}{
			"id":            product.ID,
			"user_id":       product.UserID,
			"catagory_id":   product.CatagoryID,
			"product_title": product.ProductTitle,
			"product_desc":  product.ProductDesc,
			"price":         product.Price,
			"stock":         product.Stock,
			"url_product":   product.UrlProduct,
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success update data product", responseProduct))
	}
}

func (uh *ProductHandler) DeleteProductHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

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
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success delete product"))
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

		responseProduct := map[string]interface{}{
			"id":            product.ID,
			"user_id":       product.UserID,
			"catagory_id":   product.CatagoryID,
			"product_title": product.ProductTitle,
			"product_desc":  product.ProductDesc,
			"price":         product.Price,
			"stock":         product.Stock,
			"url_product":   product.UrlProduct,
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get product by id", responseProduct))
	}
}

func (uh *ProductHandler) GetAllProductUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		products, err := uh.productUseCase.GetAllProductUser(uint(idToken))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}

		responseProducts := []map[string]interface{}{}
		for i := 0; i < len(products); i++ {
			response := map[string]interface{}{
				"id":            products[i].ID,
				"user_id":       products[i].UserID,
				"catagory_id":   products[i].CatagoryID,
				"product_title": products[i].ProductTitle,
				"product_desc":  products[i].ProductDesc,
				"price":         products[i].Price,
				"stock":         products[i].Stock,
				"url_product":   products[i].UrlProduct,
			}
			responseProducts = append(responseProducts, response)
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all tasks", responseProducts))
	}
}
