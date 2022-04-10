package user

import (
	"group-project/limamart/delivery/helper"
	_middlewares "group-project/limamart/delivery/middlewares"
	_userUseCase "group-project/limamart/usecase/users"
	"net/http"
	"strconv"

	_entities "group-project/limamart/entities"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase _userUseCase.UserUseCaseInterface
}

func NewUserHandler(u _userUseCase.UserUseCaseInterface) UserHandler {
	return UserHandler{
		userUseCase: u,
	}
}

func (uh *UserHandler) CreateUserHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		var param _entities.User

		errBind := c.Bind(&param)
		if errBind != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errBind.Error()))
		}
		_, err := uh.userUseCase.CreateUser(param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success create user"))
	}
}

func (uh *UserHandler) UpdateUserHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		var param _entities.User
		id, _ := strconv.Atoi(c.Param("id"))

		if idToken != id {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		err := c.Bind(&param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		users, rows, err := uh.userUseCase.UpdateUser(id, param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}

		responseUser := map[string]interface{}{
			"id":           users.ID,
			"name":         users.Name,
			"email":        users.Email,
			"phone_number": users.PhoneNumber,
			"address":      users.Address,
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success update data", responseUser))
	}
}

func (uh *UserHandler) DeleteUserHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		id, _ := strconv.Atoi(c.Param("id"))

		if idToken != id {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		err := uh.userUseCase.DeleteUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success delete user", err))
	}
}

func (uh *UserHandler) GetUserByIdHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		user, rows, err := uh.userUseCase.GetUserById(idToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}

		responseUser := map[string]interface{}{
			"id":           user.ID,
			"name":         user.Name,
			"email":        user.Email,
			"phone_number": user.PhoneNumber,
			"address":      user.Address,
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get user by id", responseUser))
	}
}
