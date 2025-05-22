package controller

import (
	"go-api/model"
	"go-api/service"
	"go-api/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type StockController struct {
	Service *service.StockService
}

func NewStockController(service *service.StockService) *StockController {
	return &StockController{Service: service}
}

func (c *StockController) GetStockByWarehouseID(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Invalid Warehouse ID: %v", err)
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid warehouse ID")
	}

	stocks, err := c.Service.GetStockByWarehouseID(id)
	if err != nil {
		log.Printf("Error GetStockByWarehouseID: %v", err)
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to retrieve stock by warehouse")
	}

	return utils.Respond(ctx, http.StatusOK, stocks, "Success")
}

func (c *StockController) GetStockByWarehouseProductID(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		log.Printf("Invalid Warehouse ID: %v", err)
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid warehouse ID")
	}

	productId := ctx.Param("product_id")
	product_id, err := strconv.Atoi(productId)

	if err != nil {
		log.Printf("Invalid Product ID: %v", err)
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid Product ID")
	}

	stocks, err := c.Service.GetStockByWarehouseProductID(id, product_id)
	if err != nil {
		log.Printf("Error GetStockByWarehouseProductID: %v", err)
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to retrieve stock by warehouse and product")
	}

	return utils.Respond(ctx, http.StatusOK, stocks, "Success")
}

func (c *StockController) CreateStock(ctx echo.Context) error {

	idWarehouseStr := ctx.Param("id")
	idWarehouse, err := strconv.Atoi(idWarehouseStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid warehouse ID"})
	}

	var stock model.WarehouseStock
	if err := ctx.Bind(&stock); err != nil {
		log.Printf("Error binding CreateStock: %v", err)
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid request")
	}

	stock.WarehouseID = idWarehouse

	result, err := c.Service.CreateStock(stock)
	if err != nil {
		log.Printf("Error CreateStock: %v", err)
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to create stock")
	}

	return utils.Respond(ctx, http.StatusCreated, result, "Stock created successfully")
}

func (c *StockController) UpdateStock(ctx echo.Context) error {
	idWarehouseStr := ctx.Param("id")
	idWarehouse, err := strconv.Atoi(idWarehouseStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid warehouse ID"})
	}

	stockIDStr := ctx.Param("stock_id")
	stockID, err := strconv.Atoi(stockIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid stock ID"})
	}

	productIDStr := ctx.Param("product_id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	var stock model.WarehouseStock
	if err := ctx.Bind(&stock); err != nil {
		log.Printf("Error binding UpdateStock: %v", err)
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid request")
	}

	stock.ID = stockID
	stock.ProductID = productID
	stock.WarehouseID = idWarehouse

	result, err := c.Service.UpdateStock(stock)
	if err != nil {
		log.Printf("Error UpdateStock: %v", err)
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to update stock")
	}

	return utils.Respond(ctx, http.StatusOK, result, "Stock updated successfully")
}
