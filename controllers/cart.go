package controllers

import (
	"e-commerce/helpers"
	"e-commerce/middlewares"
	"e-commerce/models"
	"e-commerce/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CartController struct {
	cartService services.CartServiceInterface
}

func NewCartController(cartService services.CartServiceInterface) *CartController {
	return &CartController{
		cartService: cartService,
	}
}

func (cc *CartController) NewCart(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, helpers.APIResponseFailed(echo.ErrForbidden.Code, errToken.Error()))
	}

	var newCart models.NewCart
	err := c.Bind(&newCart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.APIResponseFailed(echo.ErrBadRequest.Code, err.Error()))
	}

	ctx := c.Request().Context()
	errNewCart := cc.cartService.NewCart(ctx, newCart, idToken)
	if errNewCart != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, errNewCart.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccessWithoutData(200, "success add new cart"))
}
