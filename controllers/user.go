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
		return c.JSON(http.StatusBadRequest, helpers.APIResponseFailed(echo.ErrBadRequest.Code, err.Error()))
	}

	ctx := c.Request().Context()
	errCreateUser := uc.userService.CreateUser(ctx, newUser)
	if errCreateUser != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, "failed create a new user"))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccessWithoutData(200, "Success create a new user"))
}

func (uc *UserController) GetUserById(c echo.Context) error {
	idString := c.Param("idUser")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, "Insert id_user"))
	}

	ctx := c.Request().Context()
	user, err := uc.userService.GetUserById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, "id_user not found"))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccess(200, "Success get user by id", user))
}

func (uc *UserController) GetAllUsers(c echo.Context) error {
	ctx := c.Request().Context()
	users, err := uc.userService.GetAllUsers(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, "Failed get all users"))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccess(200, "Success get all users", users))
}

func (uc *UserController) DeleteUser(c echo.Context) error {
	// idToken, errToken := middlewares.ExtractToken(c)
	// if errToken != nil {
	// 	return c.JSON(http.StatusUnauthorized, helpers.APIResponseFailed(echo.ErrBadRequest.Code, "unauthorized"))
	// }
	idString := c.Param("idUser")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, "Insert id_user"))
	}
	ctx := c.Request().Context()
	errDelete := uc.userService.DeleteUser(ctx, id)
	if errDelete != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, "Token invalid"))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccessWithoutData(200, "User deleted"))
}

func (uc *UserController) UpdateUser(c echo.Context) error {
	idString := c.Param("idUser")
	id, errToken := strconv.Atoi(idString)
	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, "insert id_user"))
	}

	var updateUser models.UserUpdate
	err := c.Bind(&updateUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.APIResponseFailed(echo.ErrBadRequest.Code, "update user failed"))
	}

	ctx := c.Request().Context()
	user, errUpdate := uc.userService.UpdateUser(ctx, updateUser, id)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(echo.ErrBadGateway.Code, "update user failed"))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccess(200, "success update user", user))
}
