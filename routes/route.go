package routes

import (
	"e-commerce/controllers"

	"github.com/labstack/echo/v4"
)

func UserPath(e *echo.Echo, uc *controllers.UserController) {
	e.POST("/users", uc.CreateUser)
	e.GET("/users/:idUser", uc.GetUserById)
}
