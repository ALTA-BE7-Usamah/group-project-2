package routes

import (
	_authHandler "group-project/limamart/delivery/handler/auth"
	_cartHandler "group-project/limamart/delivery/handler/cart"
	_catagoryHandler "group-project/limamart/delivery/handler/catagory"
	_orderHandler "group-project/limamart/delivery/handler/order"
	_productHandler "group-project/limamart/delivery/handler/product"
	_userHandler "group-project/limamart/delivery/handler/users"
	_middlewares "group-project/limamart/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/auth", ah.LoginHandler())
}

func RegisterUserPath(e *echo.Echo, uh _userHandler.UserHandler) {
	e.GET("/users", uh.GetUserByIdHandler(), _middlewares.JWTMiddleware())
	e.POST("/users", uh.CreateUserHandler())
	e.PUT("/users/:id", uh.UpdateUserHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/users/:id", uh.DeleteUserHandler(), _middlewares.JWTMiddleware())
}

func RegisterProductPath(e *echo.Echo, uh _productHandler.ProductHandler) {
	e.GET("/products", uh.GetAllHandler())
	e.GET("/products/:id", uh.GetProductByIdHandler())
	e.GET("products/users", uh.GetAllProductUserHandler(), _middlewares.JWTMiddleware())
	e.POST("/products", uh.CreateProductHandler(), _middlewares.JWTMiddleware())
	e.PUT("/products/:id", uh.UpdateProductHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/products/:id", uh.DeleteProductHandler(), _middlewares.JWTMiddleware())
}

func RegisterCartPath(e *echo.Echo, uh _cartHandler.CartHandler) {
	e.GET("/cart", uh.GetAllHandler(), _middlewares.JWTMiddleware())
	e.POST("/cart", uh.CreateCartHandler(), _middlewares.JWTMiddleware())
	e.PUT("/cart/:id", uh.UpdateCartHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/cart/:id", uh.DeleteCartHandler(), _middlewares.JWTMiddleware())
}

func RegisterCatagoryPath(e *echo.Echo, uh _catagoryHandler.CatagoryHandler) {
	e.GET("/catagories", uh.GetAllCatagoryHandler(), _middlewares.JWTMiddleware())
}

func RegisterOrderPath(e *echo.Echo, uh _orderHandler.OrderHandler) {
	e.GET("/order", uh.GetAllOrdersHandler(), _middlewares.JWTMiddleware())
	e.POST("/order", uh.CreateOrderHandler(), _middlewares.JWTMiddleware())
}
