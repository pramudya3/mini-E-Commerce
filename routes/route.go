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

func CategoryPath(e *echo.Echo, cc *controllers.CategoryController) {
	e.POST("/categories", cc.CreateCategory, middlewares.JWTMiddleware())
	e.GET("/categories", cc.GetCategory)
	e.GET("/categories/:id", cc.GetCategoryById)
	e.DELETE("/categories/:id", cc.DeleteCategory, middlewares.JWTMiddleware())
	e.PUT("/categories/:id", cc.UpdateCategory, middlewares.JWTMiddleware())
}

func ProductPath(e *echo.Echo, pc *controllers.ProductController) {
	e.POST("/products", pc.NewProduct, middlewares.JWTMiddleware())
	e.GET("/products/:id", pc.GetProductById)
	e.GET("/products", pc.GetAllProducts)
	e.DELETE("/products/:id", pc.DeleteProduct, middlewares.JWTMiddleware())
	e.PUT("/products/:id", pc.UpdateProduct, middlewares.JWTMiddleware())
}

func CartPath(e *echo.Echo, cc *controllers.CartController) {
	e.POST("/carts", cc.NewCart, middlewares.JWTMiddleware())
}

func OrderPath(e *echo.Echo, oc *controllers.OrderController) {
	e.POST("/orders", oc.NewOrder, middlewares.JWTMiddleware())
}
