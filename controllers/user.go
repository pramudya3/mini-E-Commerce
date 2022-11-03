package controllers

import (
	"e-commerce/helpers"
	"e-commerce/models"
	"e-commerce/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService services.UserServiceInterface
}

func NewUserController(userService services.UserServiceInterface) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) CreateUser(c echo.Context) error {
	var newUser models.User
	err := c.Bind(&newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.APIResponseFailed(echo.ErrBadRequest.Code, "Failed to access database"))
	}

	ctx := c.Request().Context()
	errCreateUser := uc.userService.CreateUser(ctx, newUser)
	if errCreateUser != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, "Failed create a new user"))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccessWithoutData(200, "Success create a new user"))
}

func (uc *UserController) GetUserById(c echo.Context) error {
	idString := c.Param("idUser")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(400, "Id not recognize"))
	}

	ctx := c.Request().Context()
	user, err := uc.userService.GetUserById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(400, "Id not recognize"))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccess(200, "Success get user by id", user))
}
