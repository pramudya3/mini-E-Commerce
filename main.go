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

	// users
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	// login
	loginRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(loginRepository)
	authController := controllers.NewAuthController(authService)

	// category
	catRepository := repositories.NewCategoryRepository(db)
	catService := services.NewCategoryService(catRepository)
	catController := controllers.NewCategoryController(catService)

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productCtrl := controllers.NewProductController(productService)

	// route
	e := echo.New()
	routes.UserPath(e, userController)
	routes.LoginPath(e, authController)
	routes.CategoryPath(e, catController)
	routes.ProductPath(e, productCtrl)

	// start server
	log.Fatal(e.Start(":1234"))
}
