package main

import (
	"e-commerce/controllers"
	"e-commerce/databases"
	"e-commerce/repositories"
	"e-commerce/routes"
	"e-commerce/services"
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	db := databases.GetCoonectMysql()
	defer db.Close()

	// users
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	// login
	loginRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(loginRepository)
	authController := controllers.NewAuthController(authService)

	// route
	e := echo.New()
	routes.UserPath(e, userController)
	routes.LoginPath(e, authController)

	fmt.Println(time.Now())

	// start server
	log.Fatal(e.Start(":1234"))
}
