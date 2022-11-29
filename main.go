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
	authCtrl := controllers.NewAuthController(authService)

	// category
	catRepository := repositories.NewCategoryRepository(db)
	catService := services.NewCategoryService(catRepository)
	catCtrl := controllers.NewCategoryController(catService)

	// product
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productCtrl := controllers.NewProductController(productService)

	// order
	orderRepo := repositories.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderCtrl := controllers.NewOrderController(orderService)

	// cart
	cartRepo := repositories.NewCartRepositor(db)
	cartService := services.NewCartService(cartRepo)
	cartCtrl := controllers.NewCartController(cartService)

	// route
	e := echo.New()
	routes.UserPath(e, userController)
	routes.LoginPath(e, authCtrl)
	routes.CategoryPath(e, catCtrl)
	routes.ProductPath(e, productCtrl)
	routes.OrderPath(e, orderCtrl)
	routes.CartPath(e, cartCtrl)

	// start server
	log.Fatal(e.Start(":1234"))
}
