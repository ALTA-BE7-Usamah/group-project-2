package catagory

import (
	"group-project/limamart/delivery/helper"
	_catagoryUseCase "group-project/limamart/usecase/catagory"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CatagoryHandler struct {
	catagoryUseCase _catagoryUseCase.CatagoryUseCaseInterface
}

func NewCatagoryHandler(c _catagoryUseCase.CatagoryUseCaseInterface) CatagoryHandler {
	return CatagoryHandler{
		catagoryUseCase: c,
	}
}

func (uh *CatagoryHandler) GetAllCatagoryHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		catagory, err := uh.catagoryUseCase.GetAllCatagory()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all catagories", catagory))
	}
}
