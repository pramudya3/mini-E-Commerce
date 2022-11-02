package main

import (
	"e-commerce/controllers"
	"e-commerce/databases"
	"e-commerce/repositories"
	"e-commerce/routes"
	"e-commerce/services"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	db := databases.GetCoonectMysql()
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	e := echo.New()
	routes.UserPath(e, userController)

	log.Fatal(e.Start(":1234"))
}
