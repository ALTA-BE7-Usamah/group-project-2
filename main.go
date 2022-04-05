package main

import (
	"fmt"
	"group-project/limamart/configs"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_authHandler "group-project/limamart/delivery/handler/auth"
	_authRepository "group-project/limamart/repository/auth"
	_authUseCase "group-project/limamart/usecase/auth"

	_userHandler "group-project/limamart/delivery/handler/users"
	_userRepository "group-project/limamart/repository/users"
	_userUseCase "group-project/limamart/usecase/users"

	_productHandler "group-project/limamart/delivery/handler/product"
	_productRepository "group-project/limamart/repository/product"
	_productUseCase "group-project/limamart/usecase/product"

	_middlewares "group-project/limamart/delivery/middlewares"
	_routes "group-project/limamart/delivery/routes"
	_utils "group-project/limamart/utils"
)

func main() {
	config := configs.GetConfig()
	db := _utils.InitDB(config)

	authRepo := _authRepository.NewAuthRepository(db)
	authUseCase := _authUseCase.NewAuthUseCase(authRepo)
	authHandler := _authHandler.NewAuthHandler(authUseCase)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)

	productRepo := _productRepository.NewProductRepository(db)
	productUseCase := _productUseCase.NewProductUseCase(productRepo)
	productHandler := _productHandler.NewProductHandler(productUseCase)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(_middlewares.CustomLogger())

	_routes.RegisterAuthPath(e, authHandler)
	_routes.RegisterUserPath(e, userHandler)
	_routes.RegisterProductPath(e, productHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
