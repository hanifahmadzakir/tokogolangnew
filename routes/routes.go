package routes

import (
	"tokogolangnew/handler"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, userHandler *handler.UserHandler, productHandler *handler.ProductHandler, cartHandler *handler.CartHandler) {
	e.POST("/users/create", userHandler.CreateUser)
	e.GET("/users", userHandler.GetAllUser)
	e.GET("/users/:id", userHandler.GetUserByID)
	e.PUT("/users/update/:id", userHandler.UpdateUser)
	e.DELETE("/users/delete/:id", userHandler.DeleteUser)

	e.POST("/products/create", productHandler.CreateProduct)
	e.GET("/products", productHandler.GetAllProduct)
	e.GET("/products/:id", productHandler.GetProductByID)
	e.GET("/products/:category", productHandler.GetProductByCategory)
	e.GET("/products/:priceMin/:priceMax", productHandler.GetProductByPrice)
	e.PUT("/products/update/:id", productHandler.UpdateProduct)
	e.DELETE("/products/delete/:id", productHandler.DeleteProduct)

	e.POST("/carts/create", cartHandler.CreateCart)
	e.GET("/carts", cartHandler.SelectAllCart)
	e.GET("/carts/:id", cartHandler.GetCartByID)
	e.PUT("/carts/update/:id", cartHandler.UpdateCart)
	e.DELETE("/carts/delete/:id", cartHandler.DeleteCart)
	e.PATCH("/carts/payment/:id", cartHandler.Payment)

}
