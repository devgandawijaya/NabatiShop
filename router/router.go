package router

import (
	"go-api/controller"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo,
	shopController *controller.ShopController,
	warehouseController *controller.WarehouseController,
	stockController *controller.StockController,
	transferController *controller.TransferController,
) {

	e.POST("/shops", shopController.CreateShop)
	e.GET("/shops", shopController.GetAllShop)
	e.PUT("/shops/:id", shopController.UpdateShop)
	e.DELETE("/shops/:id", shopController.DeleteShop)
	e.POST("/warehouses", warehouseController.CreateWarehouses)
	e.GET("/warehouses", warehouseController.GetAllWarehouses)
	e.PUT("/warehouses/:id", warehouseController.UpdateWarehouses)
	e.DELETE("/warehouses/:id", warehouseController.DeleteWarehouses)
	e.GET("/warehouses/:id/stocks", stockController.GetStockByWarehouseID)
	e.GET("/warehouses/:id/stocks/:product_id", stockController.GetStockByWarehouseProductID)
	e.PUT("/warehouses/:id/stocks/:stock_id/product/:product_id", stockController.UpdateStock)
	e.POST("/warehouses/:id/stocks", stockController.CreateStock)
	e.POST("/transfers", transferController.CreateTransfer)
	e.GET("/transfers", transferController.GetAllTransfer)
}
