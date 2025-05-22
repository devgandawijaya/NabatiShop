package main

import (
	"go-api/config"
	"go-api/controller"
	"go-api/repository"
	"go-api/router"
	"go-api/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.InitDB()
	e := echo.New()
	e.Use(middleware.CORS())
	shopRepo := repository.NewShopRepository(config.DB)
	shopService := service.NewShopService(shopRepo)
	shopController := controller.NewShopController(shopService)

	warehouseRepo := repository.NewWarehouseRepository(config.DB)
	warehouseService := service.NewWarehouseService(warehouseRepo)
	warehouseController := controller.NewWarehouseController(warehouseService)

	stockRepo := repository.NewStockRepository(config.DB)
	stockService := service.NewStockService(stockRepo)
	stockController := controller.NewStockController(stockService)

	trasnferRepo := repository.NewTransferRepository(config.DB)
	trasnferService := service.NewTransferService(trasnferRepo)
	trasnferController := controller.NewTransferController(trasnferService)

	router.InitRouter(e, shopController, warehouseController, stockController, trasnferController)

	e.Logger.Fatal(e.Start(":8080"))
}
