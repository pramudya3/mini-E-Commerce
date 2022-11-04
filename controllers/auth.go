package controllers

import (
	"e-commerce/helpers"
	"e-commerce/models"
	"e-commerce/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authService services.AuthServiceInterface
}

func NewAuthController(authService services.AuthServiceInterface) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) Login(c echo.Context) error {
	var userLogin models.LoginRequest
	err := c.Bind(&userLogin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.APIResponseFailed(echo.ErrBadRequest.Code, err.Error()))
	}

	ctx := c.Request().Context()
	loginResponse, err := ac.authService.Login(ctx, userLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccess(200, "success login", loginResponse))
}
