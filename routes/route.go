package routes

import (
	"e-commerce/controllers"
	"e-commerce/middlewares"

	"github.com/labstack/echo/v4"
)

func UserPath(e *echo.Echo, uc *controllers.UserController) {
	e.POST("/users", uc.CreateUser)
	e.GET("/users/:idUser", uc.GetUserById)
	e.GET("/users", uc.GetAllUsers)
	e.DELETE("/users", uc.DeleteUser, middlewares.JWTMiddleware())
	e.PUT("/users", uc.UpdateUser, middlewares.JWTMiddleware())
}

func LoginPath(e *echo.Echo, ac *controllers.AuthController) {
	e.POST("/login", ac.Login)
}
