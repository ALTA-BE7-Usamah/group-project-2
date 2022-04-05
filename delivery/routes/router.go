package routes

import (
	_authHandler "group-project/limamart/delivery/handler/auth"
	_userHandler "group-project/limamart/delivery/handler/users"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/auth", ah.LoginHandler())
}

func RegisterUserPath(e *echo.Echo, uh _userHandler.UserHandler) {
	e.POST("/users", uh.CreateUserHandler())
}