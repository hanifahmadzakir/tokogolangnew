package main

import (
	"tokogolangnew/config"
	// "tokogolangnew/docs"
	"tokogolangnew/handler"
	"tokogolangnew/repository"
	"tokogolangnew/routes"
	"tokogolangnew/usecase"

	"github.com/labstack/echo/v4"
)

func main () {
	config.Database()
	config.Migrate()

	e := echo.New()

	userRepository := repository.NewUserRepository(config.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	productRepository := repository.NewProductRepository(config.DB)
	productUsecase := usecase.NewProductUsecase(productRepository)
	productHandler := handler.NewProductHandler(productUsecase)

	cartRepository := repository.NewCartRepository(config.DB)
	cartUsecase := usecase.NewCartUsecase(cartRepository)
	cartHandler := handler.NewCartHandler(cartUsecase)

	routes.Routes(e, userHandler, productHandler, cartHandler)

	e.Logger.Fatal(e.Start(":9000"))
	
}