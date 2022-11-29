package controllers

import (
	"e-commerce/helpers"
	"e-commerce/middlewares"
	"e-commerce/models"
	"e-commerce/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	orderService services.OrderServiceInterface
}

func NewOrderController(orderService services.OrderServiceInterface) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

func (oc *OrderController) NewOrder(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, helpers.APIResponseFailed(echo.ErrForbidden.Code, errToken.Error()))
	}

	var newOrder models.NewOrder
	err := c.Bind(&newOrder)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.APIResponseFailed(echo.ErrBadRequest.Code, err.Error()))
	}

	ctx := c.Request().Context()
	errNewOrder := oc.orderService.NewOrder(ctx, newOrder, idToken)
	if errNewOrder != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, errNewOrder.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccessWithoutData(200, "success add new order"))
}
